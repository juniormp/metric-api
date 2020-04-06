package repository

import (
	"errors"

	"github.com/juniormp/metric-api/src/domain"
)

func (redisAdapter RedisAdapter) AddMetricRepository(metric domain.Metric) error {
	_, err := redisAdapter.Client.LPush(string(metric.Name), metric.ExpiredAt).Result()

	return err
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

	metrics = domain.Metrics{metricName, values}

	return metrics, err
}