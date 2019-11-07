package main

import (
	"flag"
	"os"

	"github.com/cnnrrss/nbaconsole/app"
)

var gameDate, timeZone string

func init() {
	flag.StringVar(&gameDate, "d", "", "optionally retrieve NBA scoreboard for date in YYYYMMDD format")
	flag.StringVar(&timeZone, "tz", "America/Los_Angeles", "optionally set time zone")
}

func main() {
	flag.Parse()
	var debug bool
	if os.Getenv("DEBUG") != "" {
		debug = true
	}
	nba := app.NewNBAConsole(gameDate, timeZone, debug)
	nba.Start()
}
