package repository

import (
	"errors"

	"github.com/go-redis/redis"
	"github.com/juniormp/metric-api/src/domain"
)

type MetricsRepository struct {
	Adapter RedisAdapter
}

func (metricsRepository MetricsRepository) getClient() *redis.Client {
	return metricsRepository.Adapter.Client
}

func (repository MetricsRepository) AddMetricRepository(metric domain.Metric) error {
	client := repository.getClient()
	_, err := client.LPush(metric.Name, metric.ExpiredAt).Result()

	return err
}

func (repository MetricsRepository) ListKeysRepository() []string {
	client := repository.getClient()
	result, err := client.Keys("*").Result()

	if err != nil {
		errors.New("Error while listing keys")
	}

	return result
}

func (repository MetricsRepository) DeleteMetricRepository(metricName string, expiredAt string) (int64, error) {
	client := repository.getClient()
	result, err := client.LRem(metricName, 1, expiredAt).Result()

	if err != nil {
		return result, errors.New("Error while deleting metric")
	}

	return result, err
}

func (repository MetricsRepository) ListMetricsRepository(metricName string) (domain.Metrics, error) {
	var metrics domain.Metrics
	client := repository.getClient()
	qtdMetrics, err := client.LLen(metricName).Result()

	if err != nil {
		return metrics, errors.New("Error while getting metric list")
	}

	if qtdMetrics == 0 {
		return metrics, errors.New("Metric do not hava values")
	}

	values, err := client.LRange(metricName, 0, qtdMetrics).Result()

	if err != nil {
		return metrics, errors.New("Error while getting metric list")
	}

	metrics = domain.Metrics{Name: metricName, Values: values}

	return metrics, err
}
