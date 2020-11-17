package service

import (
	"fmt"
	"net/http"

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
func GetHTML(stuID, year, month string) (*http.Response, error) {
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

	return res, nil
}

// GetInfo ... 解析响应体和正则表达式匹配
func GetInfo(resp *http.Response, stuID string) {
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	strBody := string(body)
	//fmt.Println(strBody)

	reg := regexp.MustCompile("<a  class=\"w\"><h2>时间：(.*?) (.*?)</h2><span>消费：(.*?)元，</span><span>地点：后勤集团/饮食中心/(.*?)/.*?/(.*?)  </span>")
	result := reg.FindAllStringSubmatch(strBody, -1)

	reg = regexp.MustCompile("<a  class=\"w\"><h2>时间：(.*?) (.*?)</h2><span>消费：(.*?)元，</span><span>地点：后勤集团/商贸中心/(.*?)/(.*?)  </span>")
	result2 := reg.FindAllStringSubmatch(strBody, -1)

	/*
		reg4 := regexp.MustCompile("<span>地点：后勤集团/饮食中心/.*?/.*?/(.*?)</span>|<span>地点：后勤集团/商贸中心/(.*?)/.*?</span>|<span>地点：后勤集团/饮食中心/学子中西餐厅/(.*?)</span>|<span>地点：后勤集团/饮食中心/冷库中心/(.*?)</span>")
		result4 := reg4.FindAllStringSubmatch(strBody, -1)
	*/

	//检查数据是否正确，错误直接停掉，避免存入错误信息
	result = append(result, result2...)
	fmt.Println(len(result))
	if len(result) == 0 {
		log.Fatal("wrong")
	}

	var data []model.StudentsModel

	for i := 0; i < len(result); i++ {
		if len(result[i]) != 6 {
			log.Fatal("wrong")
		}
		data = append(data, model.StudentsModel{
			StuID:      stuID,
			Date:       result[i][1],
			Time:       result[i][2],
			Cost:       result[i][3],
			Restaurant: result[i][4],
			Place:      result[i][5],
		})

	}
	fmt.Println(model.DB.Self.Create(data).Error)
}

/*func Crawler(stuId, year, month string) {
	res, err := GetHtml(stuId, year, month)
	if err != nil {
		// log.Error("GetHtml Error",)
	}
}*/
