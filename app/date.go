package nbaconsole

import "time"

const (
	// YYYYMMDD is the NBA API Date format
	YYYYMMDD = "2006-01-02"
)

func currentDate() string {
	return time.Now().AddDate(0, 0, -1).Format(YYYYMMDD)
}
