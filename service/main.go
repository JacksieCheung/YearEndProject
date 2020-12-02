package main

import (
	"YearEndProject/service/config"
	"YearEndProject/service/log"
	"YearEndProject/service/model"

	"github.com/JacksieCheung/YearEndProject/service/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error

	defer log.SyncLogger()

	// init config
	err = config.Init("./conf/config.yaml", "data")
	if err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	g := gin.Default()

	g.GET("/data", handler.Data)
	g.Run("localhost:8899")
}
