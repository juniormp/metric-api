package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juniormp/metric-api/src/application/usecases"
	"github.com/juniormp/metric-api/src/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.RedisHandler)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "live!"})
	})

	router.POST("/metric/:name", usecases.AddMetric)
	router.GET("/metric/:name", usecases.MetricReport)
	router.GET("/clean-metrics", usecases.CleanMetrics)

	return router
}
