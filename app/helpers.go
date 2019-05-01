package app

import (
	"fmt"
	"time"
)

const (
	// YYYYMMDD is the NBA API Date format
	YYYYMMDD = "20060102"
)

var (
	nbaMessages = []string{
		"Sorry no hoops today...",
		"Ball is life...",
		"Steph curry with the shot...",
		"Ball don't lie...",
		"Ben Simmons for 3...",
	}
)

// genericParams returns basic parameters to include in an API request
func genericParams(date string) map[string]string {
	params := map[string]string{
		"DayOffset": "0",
		"LeagueID":  "00",
		"gameDate":  date,
	}
	return params
}

// currentDate returns the current unix time date in YYYYMMDD format
func currentDate() string {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	return time.Now().In(loc).AddDate(0, 0, 0).Format(YYYYMMDD)
}

func toHumanDate(date string) string {
	t, _ := time.Parse(YYYYMMDD, date)
	return t.Format("Mon Jan 02 2006")
}

func toHumanTime(t time.Time) string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}
