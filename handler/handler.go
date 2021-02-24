package handler

import (
	"runtime"
	"strconv"

	"YearEndProject/log"
	"net/http"

	"YearEndProject/pkg/errno"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Response 请求响应
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse send result
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	log.Info(message)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// SendBadRequest send bad request
func SendBadRequest(c *gin.Context, err error, data interface{}, cause string, source string) {
	code, message := errno.DecodeErr(err)
	log.Error(message,
		zap.String("cause", cause),
		zap.String("source", source))

	c.JSON(http.StatusBadRequest, Response{
		Code:    code,
		Message: message + ": " + cause,
		Data:    data,
	})
}

// SendError send error
func SendError(c *gin.Context, err error, data interface{}, cause string, source string) {
	code, message := errno.DecodeErr(err)
	log.Error(message,
		zap.String("cause", cause),
		zap.String("source", source))

	c.JSON(http.StatusInternalServerError, Response{
		Code:    code,
		Message: message + ": " + cause,
		Data:    data,
	})
}

// GetLine get line information
func GetLine() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "YearEndProject/handler/handler.go:63"
	}
	return file + ":" + strconv.Itoa(line)
}
