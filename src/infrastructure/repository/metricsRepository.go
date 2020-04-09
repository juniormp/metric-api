package repository

import (
	"errors"

	"github.com/juniormp/metric-api/src/domain"
)

type MetricsRepository struct {
	Adapter RedisAdapter
}

func (repository MetricsRepository) AddMetricRepository(metric domain.Metric) error {
	_, err := repository.Adapter.Client.LPush(metric.Name, metric.ExpiredAt).Result()

	return err
}

func (redisAdapter RedisAdapter) ListKeysRepository() []string {
	result, err := redisAdapter.Client.Keys("*").Result()

	if err != nil {
		errors.New("Error while listing keys")
	}

	return result
}

func (redisAdapter RedisAdapter) DeleteMetricRepository(metricName string, expiredAt string) (int64, error) {
	result, err := redisAdapter.Client.LRem(metricName, 1, expiredAt).Result()

	if err != nil {
		return result, errors.New("Error while deleting metric")
	}

	return result, err
}

func (redisAdapter RedisAdapter) ListMetricsRepository(metricName string) (domain.Metrics, error) {
	var metrics domain.Metrics
	qtdMetrics, err := redisAdapter.Client.LLen(metricName).Result()

	if err != nil {
		return metrics, errors.New("Error while getting metric list")
	}

	if qtdMetrics == 0 {
		return metrics, errors.New("Metric do not hava values")
	}

	values, err := redisAdapter.Client.LRange(metricName, 0, qtdMetrics).Result()

	if err != nil {
		return metrics, errors.New("Error while getting metric list")
	}

	metrics = domain.Metrics{Name: metricName, Values: values}

	return metrics, err
}
