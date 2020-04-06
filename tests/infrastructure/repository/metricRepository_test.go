package infrastructure_test

import (
	"os"
	"testing"

	"github.com/go-redis/redis"
	"github.com/juniormp/metric-api/src/domain"
	"github.com/juniormp/metric-api/src/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreatesAMetric(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_SERVER")})
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
