package app

import "time"

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
	return time.Now().AddDate(0, 0, -1).Format(YYYYMMDD)
}

func toHumanDateTime(date string) string {
	t, _ := time.Parse(YYYYMMDD, date)
	return t.Format("Mon Jan _2 15:04:05 2006")
}
