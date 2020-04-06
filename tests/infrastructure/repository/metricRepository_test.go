package infrastructure_test

import (
	"os"
	"testing"

	"github.com/go-redis/redis"
	"github.com/juniormp/metric-api/src/domain"
	"github.com/juniormp/metric-api/src/infrastructure/repository"
	factorytest "github.com/juniormp/metric-api/tests/factory"
	"github.com/stretchr/testify/assert"
)

func TestCreatesMetric(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_SERVER")})
	client.FlushAll()
	redisAdapter := repository.RedisAdapter{client}
	metric := domain.Metric{
		Name:      "clicks",
		Value:     20.00,
		ExpiredAt: "2020-04-05T15:04:05Z07:00_20",
	}

	err := redisAdapter.AddMetricRepository(metric)
	validate, _ := client.LIndex(metric.Name, 0).Result()

	assert.Equal(t, metric.ExpiredAt, validate, "it creates a metric into a redis list")
	assert.Empty(t, err)
	client.FlushAll()
}

func TestGetMetricList(t *testing.T) {
	metrics := domain.Metrics{
		Name: "clicks",
		Values: []string{
			"2020-04-05T15:04:05Z07:00_20",
			"2020-04-05T16:04:05Z07:00_20",
			"2020-04-05T17:04:05Z07:00_20",
		},
	}
	client := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_SERVER")})
	redisAdapter := repository.RedisAdapter{client}
	factorytest.PersistMetrics(metrics, redisAdapter)

	response, _ := redisAdapter.ListMetricsRepository("clicks")

	assert.Equal(t, metrics.Name, response.Name)
	assert.Equal(t, metrics.Values[0], response.Values[2])
	assert.Equal(t, metrics.Name, response.Name)
	assert.Equal(t, metrics.Values[1], response.Values[1])
	assert.Equal(t, metrics.Name, response.Name)
	assert.Equal(t, metrics.Values[2], response.Values[0])
	client.FlushAll()
}
