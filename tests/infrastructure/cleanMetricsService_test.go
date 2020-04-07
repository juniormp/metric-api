package infrastructure_test

import (
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/juniormp/metric-api/src/domain"
	"github.com/juniormp/metric-api/src/infrastructure"
	"github.com/juniormp/metric-api/src/infrastructure/repository"
	factorytest "github.com/juniormp/metric-api/tests/factory"
	"github.com/stretchr/testify/assert"
)

func TestCleansOldMetrics(t *testing.T) {
	validMetric := time.Now().Add(time.Hour).Format("2006-01-02T15:04:05Z07:00") + "_40"

	metrics := domain.Metrics{
		Name: "clicks",
		Values: []string{
			validMetric,
			"2020-01-07T04:23:48Z_13",
			"2020-01-07T04:23:48Z_13",
		},
	}
	client := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_SERVER")})
	client.FlushAll()
	redisAdapter := repository.RedisAdapter{client}
	factorytest.PersistMetrics(metrics, redisAdapter)

	infrastructure.CleanMetrics(redisAdapter)
	result, _ := redisAdapter.ListMetricsRepository(metrics.Name)

	assert.Equal(t, validMetric, result.Values[0])

}
