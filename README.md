# NBA Console
A minimal NBA scoreboard presented in the terminal using gocui

## Usage 
`cd path/to/nbaconsole`
`go mod`
`make clean && make build`
```
./bin/nbaconsole -d <YYYYMMDD> (optional date flag)
```

## Features
Scoreboard View: View todays games
Poll changes to scoreboard automatically using lightweight goroutine
Optionally view scoreboard on any date
Developer Productivity: Track NBA scoress while remianing in the terminal

Action|Description
------|---------------|
`quit`|Quit NBAConsole

### Roadmap
\- Expand game to full boxscore

#### Known Issues
\- Redraw scoreboard with header when terminal window changes
\- Ability to scroll scoreboard when terminal height is minified

[![Go Report Card](https://goreportcard.com/badge/github.com/connorvanderhook/nbaconsole?style=flat-square)](https://goreportcard.com/report/github.com/connorvanderhook/nbaconsole)