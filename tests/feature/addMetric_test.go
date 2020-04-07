package feature_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/bluele/go-timecop"
	"github.com/go-redis/redis"
	"github.com/juniormp/metric-api/src/infrastructure/repository"
	"github.com/juniormp/metric-api/src/web"
	"github.com/stretchr/testify/assert"
)

func performRequestAddMetric(r http.Handler, method, path string, metricValue float64) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(map[string]float64{
		"value": metricValue,
	})

	request, _ := http.NewRequest(method, path, bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	return w
}

func TestAddMetricFeature(t *testing.T) {
	metricName := "payments"
	metricValue := 20.00
	currentTime := timecop.Now()
	timeToExpire := currentTime.Add(time.Hour)
	expiredAt := timeToExpire.Format("2006-01-02T15:04:05Z07:00") + "_" + "20"

	router := web.SetupRouter()
	response := performRequestAddMetric(router, "POST", "/metric/"+metricName, metricValue)

	client := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_SERVER")})
	redisAdapter := repository.RedisAdapter{client}
	metrics, err := redisAdapter.ListMetricsRepository("payments")

	assert.Equal(t, metricName, metrics.Name)
	assert.Equal(t, expiredAt, metrics.Values[0])
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, response.Code)
	client.FlushAll()
}
