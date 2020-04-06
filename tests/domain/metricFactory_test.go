package domain_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/bluele/go-timecop"
	"github.com/juniormp/metric-api/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreatesMetric(t *testing.T) {
	name := "clicks"
	value := 20.00
	currentTime := timecop.Now()
	timeToExpire := currentTime.Add(time.Hour)
	expiredAt := timeToExpire.Format("2006-01-02T15:04:05Z07:00") + "_" + fmt.Sprintf("%.0f", value)

	response := domain.CreateMetric(name, value)

	assert.Equal(t, response.Name, name, "The two words should be the same.")
	assert.Equal(t, response.Value, value, "The two words should be the same.")
	assert.Equal(t, response.ExpiredAt, expiredAt, "The two words should be the same.")
}

func TestOnCreatingMetricValueMustBeRounded(t *testing.T) {
	name := "clicks"
	value := 27.75
	valueRounded := "28"
	currentTime := timecop.Now()
	timeToExpire := currentTime.Add(time.Hour)
	expiredAt := timeToExpire.Format("2006-01-02T15:04:05Z07:00") + "_" + valueRounded

	response := domain.CreateMetric(name, value)

	assert.Equal(t, response.ExpiredAt, expiredAt, "The two words should be the same.")
}
