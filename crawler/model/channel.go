package model

// ErrorMsg ErrChannel
type ErrorMsg struct {
	Err          error
	ChannelIndex string // 管道信息：1->1月，2->2月 ... 0->database
	StuID        string
	Result       [][]string
	Data         string // 当前数据，显示错误的是哪个信息的
}

// IndexMsg IndChannel
type IndexMsg struct {
	Month string // 月份
	StuID string // 学号
}

// RespMsg ResChannel
type RespMsg struct {
	StuID  string
	Result [][]string
}
