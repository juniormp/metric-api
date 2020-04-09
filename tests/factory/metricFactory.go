package factorytest

import (
	"github.com/go-redis/redis"
	"github.com/juniormp/metric-api/src/domain"
	"github.com/juniormp/metric-api/src/infrastructure/repository"
)

func getClient(metricsRepository repository.MetricsRepository) *redis.Client {
	return metricsRepository.Adapter.Client
}

func PersistMetric(metric domain.Metric, metricsRepository repository.MetricsRepository) {
	client := getClient(metricsRepository)
	client.LPush(metric.Name, metric.ExpiredAt).Result()
}

func PersistMetrics(metrics domain.Metrics, metricsRepository repository.MetricsRepository) {
	client := getClient(metricsRepository)
	for _, expiredAt := range metrics.Values {
		client.LPush(metrics.Name, expiredAt).Result()
	}
}
