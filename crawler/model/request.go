package model

import (
	"github.com/spf13/viper"
)

var Request *RequestInfo

type RequestInfo struct {
	Url            string
	Cookie         string
	UserAgent      string
	ContentType    string
	Accept         string
	Host           string
	AcceptEncoding string
	Connection     string
	CacheControl   string
	ContentLength  string
	PostmanToken   string
}

func (request *RequestInfo) Init() {
	Request = &RequestInfo{
		Url:            viper.GetString("web.url"),
		Cookie:         viper.GetString("web.cookie"),
		UserAgent:      viper.GetString("web.user_agent"),
		ContentType:    viper.GetString("web.content_type"),
		Accept:         "*/*",
		Host:           viper.GetString("web.host"),
		AcceptEncoding: viper.GetString("web.accept_encoding"),
		Connection:     viper.GetString("web.connection"),
		CacheControl:   viper.GetString("web.cache_control"),
		ContentLength:  viper.GetString("web.content_length"),
		PostmanToken:   viper.GetString("web.postman_token"),
	}
}
