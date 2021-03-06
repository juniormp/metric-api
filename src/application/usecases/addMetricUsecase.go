package usecases

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/juniormp/metric-api/src/domain"
	"github.com/juniormp/metric-api/src/infrastructure/repository"
)

func AddMetric(c *gin.Context) {
	var metricRequest MetricRequest
	metricRequest.Name = c.Param("name")

	if err := c.ShouldBindJSON(&metricRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := execute(metricRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "")
}

func execute(metricRequest MetricRequest) error {
	client := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_SERVER")})
	redisAdapter := repository.RedisAdapter{Client: client}
	repository := repository.MetricsRepository{Adapter: redisAdapter}

	metric := domain.CreateMetric(metricRequest.Name, metricRequest.Value)
	err := repository.AddMetricRepository(metric)

	return err
}
