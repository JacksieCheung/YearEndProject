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

// 获取 html 源码
func GetHtml(stuId, year, month string) (*http.Response, error) {
	var build strings.Builder

	// url
	url := model.Request.Url
	build.WriteString(url)
	build.WriteString(stuId)
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

	req.Header.Add("Cookie", model.Request.Cookie)
	req.Header.Add("User-Agent", model.Request.UserAgent)
	req.Header.Add("Content-Type", model.Request.ContentType)
	req.Header.Add("Accept", model.Request.Accept)
	req.Header.Add("Cache-Control", model.Request.CacheControl)
	req.Header.Add("Postman-Token", model.Request.PostmanToken)
	req.Header.Add("Host", model.Request.Host)
	req.Header.Add("Accept-Encoding", model.Request.AcceptEncoding)
	req.Header.Add("Content-Length", model.Request.ContentLength)
	req.Header.Add("Connection", model.Request.Connection)
	req.Header.Add("cache-control", model.Request.CacheControl)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Get Html error",
			zap.String("reason", err.Error()))
		return nil, err
	}

	return res, nil
}

// 解析响应体和正则表达式匹配
func GetInfo(resp *http.Response) {
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	strBody := string(body)

	reg := regexp.MustCompile(`<a  class=\"w\"><h2>时间：(.*?)</h2><span>消费：.*?</span><span>地点：后勤集团/饮食中心/[\p{Han}]+</span>|<a  class=\"w\"><h2>时间：(.*?)</h2><span>消费：.*?</span><span>地点：后勤集团/商贸中心/[\p{Han}]+</span>`)
	result := reg.FindAllString(strBody, -1)
	/*if len(result) == 0 { //判断是否为有效学号，不是就跳过
		fmt.Println("No Data")
		return
	}*/

	/*reg2 := regexp.MustCompile("<span>消费：(.*?)元，</span><span>地点：后勤集团/饮食中心/.*?</span>|<span>消费：(.*?)元，</span><span>地点：后勤集团/商贸中心/.*?</span>")
	result2 := reg2.FindAllStringSubmatch(strBody, -1)

	reg3 := regexp.MustCompile("<span>地点：后勤集团/饮食中心/(.*?)/.*?</span>|<span>地点：后勤集团/商贸中心/.*?/(.*?)</span>")
	result3 := reg3.FindAllStringSubmatch(strBody, -1)

	reg4 := regexp.MustCompile("<span>地点：后勤集团/饮食中心/.*?/.*?/(.*?)</span>|<span>地点：后勤集团/商贸中心/(.*?)/.*?</span>|<span>地点：后勤集团/饮食中心/学子中西餐厅/(.*?)</span>|<span>地点：后勤集团/饮食中心/冷库中心/(.*?)</span>")

	result4 := reg4.FindAllStringSubmatch(strBody, -1)*/

	//检查数据是否正确，错误直接停掉，避免存入错误信息
	fmt.Println(result)
}

/*func Crawler(stuId, year, month string) {
	res, err := GetHtml(stuId, year, month)
	if err != nil {
		// log.Error("GetHtml Error",)
	}
}*/
