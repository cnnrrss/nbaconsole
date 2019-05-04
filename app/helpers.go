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

func genericParams(date string) map[string]string {
	params := map[string]string{
		"DayOffset": "0",
		"LeagueID":  "00",
		"gameDate":  date,
	}
	return params
}

func currentDateYYYYMMDD(tz string) string {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc).AddDate(0, 0, 0).Format(YYYYMMDD)
}

func toHumanDate(date string) string {
	t, _ := time.Parse(YYYYMMDD, date)
	return t.Format("Mon Jan 02 2006")
}

func toHumanTime(t time.Time) string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}
