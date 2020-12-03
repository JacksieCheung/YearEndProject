package middleware

import (
	"YearEndProject/service/handler"
	"YearEndProject/service/model"
	"YearEndProject/service/pkg/auth"
	"YearEndProject/service/pkg/errno"
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	params, err := auth.MakeAccountPreflightRequest()
	if err != nil {
		handler.SendResponse(c, errno.ErrAuthParam, err.Error())
	}

	client := http.Client{
		Timeout: auth.TIMEOUT,
	}

	var req model.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		handler.SendResponse(c, errno.ErrBind, err.Error())
		c.Abort()
		return
	}

	decodePassword, err := base64.StdEncoding.DecodeString(req.Password)
	if err != nil {
		handler.SendError(c, errno.ErrDecoding, nil, err.Error(), "")
		c.Abort()
		return
	}

	if err := auth.MakeAccountRequest(req.Id, string(decodePassword), params, &client); err != nil {
		handler.SendResponse(c, errno.ErrAuthFailed, err.Error())
		c.Abort()
		return
	}

	id, err := strconv.Atoi(req.Id)
	if err != nil {
		handler.SendError(c, errno.ErrAtoi, nil, err.Error(), "")
	}
	c.Set("id", id)

	c.Next()
}
