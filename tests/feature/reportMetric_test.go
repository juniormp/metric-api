package feature_test

import (
	"encoding/json"
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

func performRequestReportMetric(r http.Handler, method, path string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	return w
}

func TestReportMetricFeature(t *testing.T) {
	metricName := "payments"
	validMetric := time.Now().Add(time.Hour).Format("2006-01-02T15:04:05Z07:00")
	metrics := domain.Metrics{
		Name: metricName,
		Values: []string{
			validMetric + "_40",
			validMetric + "_40",
			"2020-01-07T04:23:48Z_13",
			"2020-01-07T04:23:48Z_13",
		},
	}
	client := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_SERVER")})
	redisAdapter := repository.RedisAdapter{client}
	factorytest.PersistMetrics(metrics, redisAdapter)

	router := web.SetupRouter()
	response := performRequestReportMetric(router, "GET", "/metric/"+metricName)
	var body map[string]int
	err := json.Unmarshal([]byte(response.Body.String()), &body)
	value, exists := body["value"]

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, 80, value)
	client.FlushAll()
}
