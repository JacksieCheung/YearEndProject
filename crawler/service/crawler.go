package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	//"strconv"
	"io/ioutil"
	"regexp"
	"strings"

	"YearEndProject/crawler/log"
	"YearEndProject/crawler/model"

	//"github.com/spf13/viper"
	"go.uber.org/zap"
)

// GetHTML ... 获取 html 源码
func GetHTML(stuID, year, month string) ([][]string, error) {
	var build strings.Builder

	// url
	url := model.Request.URL
	build.WriteString(url)
	build.WriteString(stuID)
	url = build.String()
	build.Reset()

	// payload
	build.WriteString("year=")
	build.WriteString(year)
	build.WriteString("&month=")
	build.WriteString(month)
	payload := strings.NewReader(build.String())
	build.Reset()

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Set("Cookie", model.Request.Cookie)
	req.Header.Set("Content-Type", model.Request.ContentType)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Get Html error",
			zap.String("reason", err.Error()))
		return nil, err
	}

	result, err := GetInfo(res)
	fmt.Println(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetInfo ... 解析响应体和正则表达式匹配
func GetInfo(resp *http.Response) ([][]string, error) {
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	strBody := string(body)
	// fmt.Println(strBody)

	reg := regexp.MustCompile("<a  class=\"w\"><h2>时间：(.*?) (.*?)</h2><span>消费：(.*?)元，</span><span>地点：后勤集团/饮食中心/(.*?)/.*?/(.*?)  </span>")
	result := reg.FindAllStringSubmatch(strBody, -1)

	reg = regexp.MustCompile("<a  class=\"w\"><h2>时间：(.*?) (.*?)</h2><span>消费：(.*?)元，</span><span>地点：后勤集团/商贸中心/(.*?)/(.*?)  </span>")
	result2 := reg.FindAllStringSubmatch(strBody, -1)

	result = append(result, result2...)
	// fmt.Println(len(result))
	if len(result) == 0 {
		return nil, errors.New("result was empty")
	}

	return result, nil
}

// InsertDataBase save data
func InsertDataBase(stuID string, result [][]string) error {
	for i := 0; i < len(result); i++ {
		if len(result[i]) != 6 {
			return errors.New("bad result format")
		}
		data := model.StudentsModel{
			StuID:      stuID,
			Date:       result[i][1],
			Time:       result[i][2],
			Cost:       result[i][3],
			Restaurant: result[i][4],
			Place:      result[i][5],
		}

		data.Create()
	}

	return nil
}

// Handler concurrency handler
func Handler() {
	// config init
	months := []int{1, 9, 10, 11}
	year := "2020"  // TODO:get from config
	stuIDMax := "0" // TODO:get from config
	stuIDMin := "0" // TODO:get from config

	stuIDEnd, _ := strconv.Atoi(stuIDMax)
	stuIDStart, _ := strconv.Atoi(stuIDMin)
	mission := (stuIDEnd - stuIDStart) * len(months)

	// channels init
	ErrChannel := make(chan model.ErrorMsg, len(months)+1) // 错误管道，database 和 crawler 共用
	ResChannel := make(chan model.RespMsg, len(months)*2)  // 结果管道， crawler 和 database 通信
	IndChannel := make(chan model.IndexMsg, len(months)*2) // 结果管道， crawler 专用，用于开启 crawler 迭代

	// goroutine init
	for _, v := range months {
		go func(stuID, month, year string) {
			result, err := GetHTML(stuID, year, month)
			if err != nil {
				ErrChannel <- model.ErrorMsg{
					Err:          err,
					StuID:        stuID,
					ChannelIndex: month,
					Data:         stuID + "&" + year + "&" + month,
				}
				return
			}

			ResChannel <- model.RespMsg{
				StuID:  stuID,
				Result: result,
			}
			IndChannel <- model.IndexMsg{
				Month: month,
				StuID: stuID,
			}
			return
		}(stuIDMin, strconv.Itoa(v), year)
	}

	// handler
	for mission > 0 {
		select {
		case msg := <-IndChannel:
			stuID, _ := strconv.Atoi(msg.StuID)
			if stuID < stuIDEnd {
				stuID++
				go func(stuID, month, year string) {
					result, err := GetHTML(stuID, year, month)
					if err != nil {
						ErrChannel <- model.ErrorMsg{
							Err:          err,
							StuID:        stuID,
							ChannelIndex: month,
							Data:         stuID + "&" + year + "&" + month,
						}
						return
					}

					ResChannel <- model.RespMsg{
						StuID:  stuID,
						Result: result,
					}
					IndChannel <- model.IndexMsg{
						Month: month,
						StuID: stuID,
					}
					mission--
					return
				}(strconv.Itoa(stuID), msg.Month, year)
			}

		case msg := <-ResChannel:
			go func(stuID string, result [][]string) {
				err := InsertDataBase(stuID, result)
				if err != nil {
					ErrChannel <- model.ErrorMsg{
						Err:          err,
						ChannelIndex: "0",
						StuID:        stuID,
						Result:       result,
						Data:         "insert into database:" + msg.StuID,
					}
					return
				}
				mission--
				return
			}(msg.StuID, msg.Result)

		case msg := <-ErrChannel: // 错误处理，重新开启 goroutine 和 log
			if msg.ChannelIndex == "0" {
				// TODO:log
				ResChannel <- model.RespMsg{
					StuID:  msg.StuID,
					Result: msg.Result,
				}
			} else {
				IndChannel <- model.IndexMsg{
					Month: msg.ChannelIndex,
					StuID: msg.StuID,
				}
			}
		}
	}
}
