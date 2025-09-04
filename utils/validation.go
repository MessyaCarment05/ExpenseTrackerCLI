package utils

import "time"

func IsValidDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	if err == nil {
		return true
	} else {
		return false
	}
}