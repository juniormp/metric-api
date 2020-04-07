package factorytest

import (
	"github.com/juniormp/metric-api/src/domain"
	"github.com/juniormp/metric-api/src/infrastructure/repository"
)

func PersistMetric(metric domain.Metric, redisAdapter repository.RedisAdapter) {
	redisAdapter.Client.LPush(metric.Name, metric.ExpiredAt).Result()
}

func PersistMetrics(metrics domain.Metrics, redisAdapter repository.RedisAdapter) {
	for _, expiredAt := range metrics.Values {
		redisAdapter.Client.LPush(metrics.Name, expiredAt).Result()
	}
}
