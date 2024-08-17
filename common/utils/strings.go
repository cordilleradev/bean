package utils

import (
	"fmt"
	"strconv"
	"time"
)

// ContainsString checks if a list of strings contains a specific string.
func ContainsString(list []string, str string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// ConvertToUTCTimestamp converts a given period (days, months, years) to a UTC timestamp
// by subtracting the period from the current time.
func ConvertToUTCTimestamp(period string) (int64, error) {
	if period == "total" {
		return 0, nil
	}

	now := time.Now().UTC()

	if len(period) < 2 {
		return 0, fmt.Errorf("invalid period format: %s", period)
	}

	value, err := strconv.Atoi(period[:len(period)-1])
	if err != nil {
		return 0, fmt.Errorf("invalid period value: %s", period)
	}

	unit := period[len(period)-1]

	var timestamp time.Time
	switch unit {
	case 'd':
		timestamp = now.AddDate(0, 0, -value)
	case 'm':
		timestamp = now.AddDate(0, -value, 0)
	case 'y':
		timestamp = now.AddDate(-value, 0, 0)
	default:
		return 0, fmt.Errorf("invalid period unit: %c", unit)
	}
	timestamp = time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
	return timestamp.Unix(), nil
}
