//并发版本
package main

import (

	// "io/ioutil"

	"YearEndProject/crawler/config"
	"YearEndProject/crawler/model"
	"YearEndProject/crawler/service"
	"fmt"

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

	res, err := service.GetHTML("2019214228", "2020", "1")
	fmt.Println(err)
	err = service.InsertDataBase("2019214228", res)
	fmt.Println(err)
	// fmt.Println(resp.Body)

	// service.GetInfo(resp, "2019214228")

	// 待修改
	/*var mession = 24
	var ch = make(chan int, 13)
	var missionMap = map[int]int{
		1:  2019214227,
		2:  2019214227,
		3:  2019214227,
		4:  2019214227,
		5:  2019214227,
		6:  2019214227,
		7:  2019214227,
		8:  2019214227,
		9:  2019214227,
		10: 2019214227,
		11: 2019214227,
		12: 2019214227,
		13: 2019214228,
	}
	//初始化并发
	go func() {
		Crawler.Get(db, missionMap[1], 1)
		ch <- 1
		mession--
	}()
	go func() {
		Crawler.Get(db, missionMap[2], 2)
		ch <- 2
		mession--
	}()
	go func() {
		Crawler.Get(db, missionMap[3], 3)
		ch <- 3
		mession--
	}()
	go func() {
		Crawler.Get(db, missionMap[4], 4)
		ch <- 4
		mession--
	}()
	go func() {
		Crawler.Get(db, missionMap[5], 5)
		ch <- 5
		mession--
	}()
	go func() {
		Crawler.Get(db, missionMap[6], 6)
		ch <- 6
		mession--
	}()
	go func() {
		Crawler.Get(db, missionMap[7], 7)
		ch <- 7
		mession--
	}()
	go func() {
		Crawler.Get(db, missionMap[8], 8)
		ch <- 8
		mession--
	}()
	go func() {
		Crawler.Get(db, missionMap[9], 9)
		ch <- 9
		mession--
	}()
	go func() {
		Crawler.Get(db, missionMap[10], 10)
		ch <- 10
		mession--
	}()
	go func() {
		Crawler.Get(db, missionMap[11], 11)
		ch <- 11
		mession--
	}()
	go func() {
		Crawler.Get(db, missionMap[12], 12)
		ch <- 12
		mession--
	}()
	for mession > 0 {
		select {
		case i := <-ch:
			missionMap[i]++
			if missionMap[i] <= missionMap[13] {
				go func() {
					Crawler.Get(db, missionMap[i], i)
					ch <- i
					mession--
				}()
			} else {
				continue
			}
		}
	}
	fmt.Println("All Finished")*/
}
