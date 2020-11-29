package handler

import (
	"YearEndProject/service/log"
	"YearEndProject/service/model"
	"YearEndProject/service/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Data return result
func Data(c *gin.Context) {
	log.Info("api server function called")

	var err error

	// get id
	id := 0

	// 获取信息
	students, err := model.GetStudentList(uint32(id))
	if err != nil {
		SendBadRequest(c, errno.ErrDatabase, nil, err.Error(), GetLine())
	}

	// 并发信息

}
