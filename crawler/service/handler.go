package service

import (
	"YearEndProject/crawler/log"
	"YearEndProject/crawler/model"
	"errors"
	"fmt"
	"strconv"

	"github.com/JacksieCheung/YearEndProject/crawler/pkg/errno"
	"go.uber.org/zap"
)

// Start concurrency handler
func Start() {
	// config init
	months := []string{"9", "10", "11"}

	stuIDEnd, _ := strconv.Atoi(model.Stu.StuIDMax)
	stuIDStart, _ := strconv.Atoi(model.Stu.StuIDMin)
	mission := (stuIDEnd - stuIDStart + 1) * len(months)

	// channels init
	ErrChannel := make(chan model.ErrorMsg, len(months)+1) // 错误管道，database 和 crawler 共用
	ResChannel := make(chan model.RespMsg, len(months)*2)  // 结果管道， crawler 和 database 通信
	IndChannel := make(chan model.IndexMsg, len(months)*2) // 结果管道， crawler 专用，用于开启 crawler 迭代
	MissionChannel := make(chan int, len(months))          // mission-- 管道

	defer close(ErrChannel)
	defer close(ResChannel)
	defer close(IndChannel)
	defer close(MissionChannel)

	fmt.Println("missions: ", mission, " start")
	// goroutine init
	for _, v := range months {
		go GoGetHTML(model.Stu.StuIDMin, v, model.Stu.Year,
			ErrChannel, ResChannel, IndChannel)
	}

	// handler
	for mission > 0 {
		select {

		case msg := <-IndChannel: // 开启下一个 GetHTML
			stuID, _ := strconv.Atoi(msg.StuID)
			if stuID < stuIDEnd {
				stuID++
				go GoGetHTML(strconv.Itoa(stuID), msg.Month, model.Stu.Year,
					ErrChannel, ResChannel, IndChannel)
			}

		case msg := <-ResChannel: // 开启 InsertDataBase
			go GoInsertDatabase(msg.StuID, msg.Month, msg.Result,
				ErrChannel, MissionChannel)

		case msg := <-ErrChannel: // 错误处理，重新开启 goroutine 和 log
			if msg.ChannelIndex == "0" { // database 插入错误，重新插一遍
				ResChannel <- model.RespMsg{
					StuID:  msg.StuID,
					Result: msg.Result,
				}

				// log
				log.Error("insert database error",
					zap.String("stuID:", msg.StuID),
					zap.String("data:", msg.Data),
					zap.String("cause:", msg.Err.Error()))
				// fmt.Println(msg.Result)
			} else if !errors.Is(errno.ErrNORESULT, msg.Err) { // 请求错误，重新来一遍
				stuID, _ := strconv.Atoi(msg.StuID)
				stuID--
				IndChannel <- model.IndexMsg{
					Month: msg.ChannelIndex,
					StuID: strconv.Itoa(stuID),
				}

				// log
				log.Error("get html error",
					zap.String("stuID:", msg.StuID),
					zap.String("month:", msg.ChannelIndex),
					zap.String("data:", msg.Data))
			} else { // html 为空，跳过
				IndChannel <- model.IndexMsg{
					Month: msg.ChannelIndex,
					StuID: msg.StuID,
				}

				log.Error("get Html is empty",
					zap.String("StuID:", msg.StuID),
					zap.String("month:", msg.ChannelIndex),
					zap.String("data:", msg.Data))
				MissionChannel <- 1
			}

		case <-MissionChannel: // 处理任务列表的 case
			mission--
			fmt.Println("mission left:", mission)
		}

	}

	fmt.Println("All done!")
}

// GoGetHTML ... 开启 GETHTML 并发
func GoGetHTML(stuID, month, year string,
	ErrChannel chan<- model.ErrorMsg, ResChannel chan<- model.RespMsg, IndChannel chan<- model.IndexMsg) {

	result, err := GetHTML(stuID, year, month)
	if err != nil {
		ErrChannel <- model.ErrorMsg{
			Err:          err,
			StuID:        stuID,
			ChannelIndex: month,
			Data:         stuID + "&" + year + "&" + month,
		}
	} else {
		ResChannel <- model.RespMsg{
			StuID:  stuID,
			Month:  month,
			Result: result,
		}
		IndChannel <- model.IndexMsg{
			Month: month,
			StuID: stuID,
		}
		log.Info("GetHTML succeeded",
			zap.String("stuID:", stuID),
			zap.String("month:", month))
	}
	return
}

// GoInsertDatabase ... 开启插入数据库并发
func GoInsertDatabase(stuID string, month string, result [][]string,
	ErrChannel chan<- model.ErrorMsg, MissionChannel chan<- int) {
	err := InsertDataBase(stuID, result)
	if err != nil {
		ErrChannel <- model.ErrorMsg{
			Err:          err,
			ChannelIndex: "0",
			StuID:        stuID,
			Result:       result,
			Data:         "insert into database:" + stuID + ",error",
		}
		return
	}
	MissionChannel <- 1
	log.Info("insert database succeed",
		zap.String("stuID:", stuID),
		zap.String("month:", month))
	return
}
