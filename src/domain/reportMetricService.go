package domain

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func CreateMetricReport(metrics Metrics) int {
	var report int = 0

	for _, expiredAtData := range metrics.Values {
		date, value := splitExpiredAtData(expiredAtData)
		currentDateTime := time.Now()

		expiredAt, err := convertStringToTime(date)

		if err != nil {
			errors.New("Error converting time")
		}

		if currentDateTime.Before(expiredAt) {
			report += convertStringInteger(value)
		}
	}

	return report
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
