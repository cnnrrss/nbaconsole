package app

import (
	"time"
)

const (
	// YYYYMMDD is the NBA API Date format
	YYYYMMDD = "20060102"
)

// currentDate returns the current unix time date in YYYYMMDD format
func currentDate() string {
	return time.Now().AddDate(0, 0, -1).Format(YYYYMMDD)
}
