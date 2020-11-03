package service

import (
	//"fmt"
	"net/http"
	//"strconv"
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

/*func Crawler(stuId, year, month string) {
	res, err := GetHtml(stuId, year, month)
	if err != nil {
		// log.Error("GetHtml Error",)
	}
}*/
