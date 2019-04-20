package cmd

import (
	"flag"
	"os"
	"time"
)

var nbaDate string

func init() {
	flag.StringVar(&nbaDate, "d", "", "optionally retrieve NBA scoreboard for date in YYYYMMDD format")
	flag.Parse()
}

// NewNBAConsole loads a new context for running the app
func NewNBAConsole() *NBAConsole {
	var debug bool
	if os.Getenv("DEBUG") != "" {
		debug = true
	}

	if nbaDate == "" {
		nbaDate = currentDate()
	}

	return &NBAConsole{
		date:          nbaDate,
		forceRefresh:  make(chan bool),
		refreshTicker: time.NewTicker(30 * time.Second),
		rateLimiter:   time.Tick(10 * time.Second),
		debug:         debug,
	}
}
