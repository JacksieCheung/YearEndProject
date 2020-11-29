package main

import (
	"YearEndProject/service/config"
	"YearEndProject/service/log"
	"YearEndProject/service/model"
)

func main() {
	var err error

	defer log.SyncLogger()

	// init config
	err = config.Init("./conf/config.yaml", "DATA")
	if err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()
}
