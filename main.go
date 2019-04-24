package main

import (
	"flag"
	"os"

	"github.com/cnnrrss/nbaconsole/app"
)

var gameDate string

func init() {
	flag.StringVar(&gameDate, "d", "", "optionally retrieve NBA scoreboard for date in YYYYMMDD format")
}

func main() {
	flag.Parse()
	var debug bool
	if os.Getenv("DEBUG") != "" {
		debug = true
	}
	nba := app.NewNBAConsole(gameDate, debug)
	nba.Start()
}
