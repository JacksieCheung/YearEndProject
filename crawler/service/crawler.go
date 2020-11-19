package service

import (
	"errors"
	"fmt"
	"net/http"

	//"strconv"
	"io/ioutil"
	"regexp"
	"strings"

	"YearEndProject/crawler/log"
	"YearEndProject/crawler/model"

	//"github.com/spf13/viper"
	"github.com/JacksieCheung/YearEndProject/crawler/pkg/errno"
	"github.com/jinzhu/gorm"
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
	// fmt.Println(result)
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
		return nil, errno.ErrNORESULT
	}

	return result, nil
}

// InsertDataBase save data
func InsertDataBase(stuID string, result [][]string) error {
	var err error
	for i := 0; i < len(result); i++ {
		if len(result[i]) != 6 {
			return errors.New("bad result format")
		}
		// find record first
		err = model.GetStudentsRecord(stuID, result[i][1], result[i][2])
		if errors.Is(gorm.ErrRecordNotFound, err) { // 只有没有记录时才会插入
			data := model.StudentsModel{
				StuID:      stuID,
				Date:       result[i][1],
				Time:       result[i][2],
				Cost:       result[i][3],
				Restaurant: result[i][4],
				Place:      result[i][5],
			}

			err := data.Create()
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		} else {
			fmt.Println("Already exists")
		}
	}
	return nil
}
