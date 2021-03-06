package usecases

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/juniormp/metric-api/src/infrastructure"
	"github.com/juniormp/metric-api/src/infrastructure/repository"
)

func CleanMetrics(c *gin.Context) {
	client := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_SERVER")})
	redisAdapter := repository.RedisAdapter{client}
	metricsRepository := repository.MetricsRepository{Adapter: redisAdapter}
	infrastructure.CleanMetrics(metricsRepository)

	c.JSON(http.StatusOK, "")
}
