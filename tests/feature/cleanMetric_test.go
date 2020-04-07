package feature_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/juniormp/metric-api/src/domain"
	"github.com/juniormp/metric-api/src/infrastructure/repository"
	"github.com/juniormp/metric-api/src/web"
	factorytest "github.com/juniormp/metric-api/tests/factory"
	"github.com/stretchr/testify/assert"
)

func performRequestCleanMetric(r http.Handler, method, path string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	return w
}

func TestCleanMetricFeature(t *testing.T) {
	metricName := "payments"
	validMetric := time.Now().Add(time.Hour).Format("2006-01-02T15:04:05Z07:00") + "_40"
	metrics := domain.Metrics{
		Name: metricName,
		Values: []string{
			validMetric,
			validMetric,
			"2020-01-07T04:23:48Z_13",
			"2020-01-07T04:23:48Z_13",
		},
	}
	client := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_SERVER")})
	client.FlushAll()
	redisAdapter := repository.RedisAdapter{client}
	factorytest.PersistMetrics(metrics, redisAdapter)

	router := web.SetupRouter()
	response := performRequestCleanMetric(router, "GET", "/clean-metrics")

	result, _ := redisAdapter.ListMetricsRepository(metricName)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, validMetric, result.Values[0])
	assert.Equal(t, validMetric, result.Values[1])
	client.FlushAll()
}
