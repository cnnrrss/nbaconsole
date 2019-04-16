package main

import (
	app "github.com/connorvanderhook/nbaconsole/app"
)

func main() {
	nba := app.NewNBAConsole()
	nba.Start()
}
