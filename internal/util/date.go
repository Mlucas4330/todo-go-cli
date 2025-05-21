package util

import (
	"fmt"
	"time"
)

func ParseDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, nil
	}
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("could not parse date '%s': %w", dateStr, err)
	}
	return t, nil
}
