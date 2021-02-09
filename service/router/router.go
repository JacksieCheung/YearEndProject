package router

import (
	"YearEndProject/service/handler/data"
	"YearEndProject/service/handler/sd"
	"YearEndProject/service/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Load the handler and middleware for router.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	dataRouter := g.Group("api/v1/data")
	dataRouter.Use(middleware.AuthMiddleware)
	{
		dataRouter.POST("", data.Data)
	}

	// The health check handlers.
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
