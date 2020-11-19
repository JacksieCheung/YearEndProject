//并发版本
package main

import (
	// "io/ioutil"

	"YearEndProject/crawler/config"
	"YearEndProject/crawler/model"
	"YearEndProject/crawler/service"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/golang/goProject/model"
)

func main() {
	var err error

	// init config
	err = config.Init("./conf/config.yaml", "CRAWLER")
	if err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	// init request information
	model.Request.Init()

	// init stu information
	model.Stu.Init()

	// start service
	service.Start()
}
