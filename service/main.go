package main

import (
	"YearEndProject/service/config"
	"YearEndProject/service/log"
	"YearEndProject/service/model"
	"YearEndProject/service/router"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/JacksieCheung/YearEndProject/service/router/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func main() {
	var err error

	defer log.SyncLogger()

	// init config
	err = config.Init("./service/conf/config.yaml", "data")
	if err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	model.FifthInfo.InitFifthPage()

	g := gin.New()

	router.Load(
		// Cores
		g,

		// Middleware
		middleware.Logging(),
		middleware.RequestId(),
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.",
				zap.String("reason", err.Error()))
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Info(
		fmt.Sprintf("Start to listening the incoming requests on http address: %s", viper.GetString("addr")))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
