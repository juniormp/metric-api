package domain

import (
	"fmt"
	"math"
	"time"
)

func CreateMetric(name string, value float64) Metric {
	return Metric{
		Name:      name,
		Value:     value,
		ExpiredAt: buildExpiredAt(value),
	}
}

func buildExpiredAt(value float64) string {
	dateTime := time.Now()
	expiredAt := dateTime.Add(time.Hour)
	value = roundNearestInteger(value)

	return expiredAt.Format("2006-01-02T15:04:05Z07:00") + "_" + convertFloatString(value)
}

func roundNearestInteger(value float64) float64 {
	return math.Round(value)
}

func convertFloatString(value float64) string {
	return fmt.Sprintf("%.0f", value)
}
