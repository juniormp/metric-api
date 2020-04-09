package domain_test

import (
	"os"
	"testing"
	"time"

	"github.com/bluele/go-timecop"
	"github.com/go-redis/redis"
	"github.com/juniormp/metric-api/src/domain"
	"github.com/juniormp/metric-api/src/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func setup(client *redis.Client) domain.Metrics {
	client.FlushAll()
	currentTime := timecop.Now()
	timeToExpire := currentTime.Add(time.Hour)

	redisAdapter := repository.RedisAdapter{client}
	repositoryMetrics := repository.MetricsRepository{Adapter: redisAdapter}
	metricName := "clicks"
	expiredAt := timeToExpire.Format("2006-01-02T15:04:05Z07:00") + "_" + "20"

	redisAdapter.Client.LPush(metricName, expiredAt).Result()

	metrics, _ := repositoryMetrics.ListMetricsRepository("clicks")

	return metrics
}

func TestGetMetricReport(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_SERVER")})
	metrics := setup(client)

	response := domain.CreateMetricReport(metrics)
	assert.Equal(t, 20, response, "")

	client.FlushAll()
}
