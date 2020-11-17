package model

import (
	"github.com/spf13/viper"
)

// Request ... 全局请求变量
var Request *RequestInfo

// RequestInfo ... 请求结构
type RequestInfo struct {
	URL         string
	Cookie      string
	ContentType string
}

// Init ... 初始化请求
func (request *RequestInfo) Init() {
	Request = &RequestInfo{
		URL:         viper.GetString("web.url"),
		Cookie:      viper.GetString("web.cookie"),
		ContentType: viper.GetString("web.content_type"),
	}
}
