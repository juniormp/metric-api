package usecases

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/juniormp/metric-api/src/domain"
	"github.com/juniormp/metric-api/src/infrastructure/repository"
)

func MetricReport(c *gin.Context) {
	metricName := c.Param("name")

	report := executeMetricReport(metricName)

	c.JSON(http.StatusOK, gin.H{
		"value": report,
	})
}

func executeMetricReport(metricName string) int {
	client := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_SERVER")})
	redisAdapter := repository.RedisAdapter{client}
	metricsRepository := repository.MetricsRepository{Adapter: redisAdapter}
	metrics, err := metricsRepository.ListMetricsRepository(metricName)

	if err != nil {
		return 0
	}

	report := domain.CreateMetricReport(metrics)

	return report
}
