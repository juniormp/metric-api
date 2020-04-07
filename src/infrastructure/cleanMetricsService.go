package infrastructure

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/juniormp/metric-api/src/infrastructure/repository"
)

func CleanMetrics(redis repository.RedisAdapter) {
	metricKeys := redis.ListKeysRepository()

	for _, key := range metricKeys {
		metrics, _ := redis.ListMetricsRepository(key)
		for _, expiredAtData := range metrics.Values {

			date, _ := splitExpiredAtData(expiredAtData)
			currentDateTime := time.Now()
			expiredAt, err := convertStringToTime(date)

			if err != nil {
				errors.New("Error converting time")
			}

			if currentDateTime.After(expiredAt) {
				redis.DeleteMetricRepository(key, expiredAtData)
			}
		}
	}
}

func splitExpiredAtData(expiredAt string) (string, string) {
	dateAndValue := strings.Split(expiredAt, "_")

	return dateAndValue[0], dateAndValue[1]
}

func convertStringToTime(date string) (time.Time, error) {
	expiredAt, err := time.Parse(time.RFC3339, date)

	return expiredAt, err
}

func convertStringInteger(valueString string) int {
	value, err := strconv.Atoi(valueString)

	if err != nil {
		errors.New("Error while converting string to integer")
	}

	return value
}
