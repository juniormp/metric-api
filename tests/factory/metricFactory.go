package factorytest

import (
	"github.com/juniormp/metric-api/src/domain"
	"github.com/juniormp/metric-api/src/infrastructure/repository"
)

func PersistMetrics(metrics domain.Metrics, redisAdapter repository.RedisAdapter) {
	for _, expiredAt := range metrics.Values {
		redisAdapter.Client.LPush(metrics.Name, expiredAt).Result()
	}
}
