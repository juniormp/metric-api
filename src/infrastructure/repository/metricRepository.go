package repository

import (
	"github.com/juniormp/metric-api/src/domain"
)

func (redisAdapter RedisAdapter) AddMetricRepository(metric domain.Metric) error {
	_, err := redisAdapter.Client.LPush(string(metric.Name), metric.ExpiredAt).Result()

	return err
}
