package cmd

import (
	"flag"

	nba "github.com/connorvanderhook/nbaconsole/app"
)

// Run the app
func Run() {
	var nbaAPIKey string
	flag.StringVar(&nbaAPIKey, "nba-api-key", "", "Api Key for Retrieving data from NBA.com")
	nba.NewNBAConsole(&nba.Config{
		APIKey: nbaAPIKey,
	}).Start()
}
