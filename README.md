# NBA Console

A minimal NBA scoreboard presented in the terminal using gocui

## Usage

`git clone git@github.com:cnnrrss/nbaconsole`

`cd path/to/nbaconsole`

`go mod`

`make clean && make build`

```bash
./bin/nbaconsole \
-d <YYYYMMDD> # (optional date flag)
```

### Controls

Screen|Action|Description
------|---------------|-------------|
All|`Ctrl+q`, `Ctrl+c`|Quit NBAConsole|
All|`Up`|Scroll up the current views' list|
All|`Down`|Scroll down the current views' list|
Scoreboard|`enter`|Expand the currently highlighted games' boxscore|
BoxScore|`Ctrl+t`|View team stats for currently selected games' boxscore|


## Features
- [x] **Scoreboard View**: View summary info of all current date games (start time, current score, final score, etc.)
- [x] **Background Refresh**: Poll changes to scoreboard automatically using lightweight background goroutine
- [x] **Specify Date**: Optionally view scoreboard on any date
    - [x] Pass in a flag `-d 20190310` with a value in `YYYYMMDD` format to view games from a specific date
- [x] **Developer Productivity**: Track NBA scoress while you continue to live in the terminal
- [ ] **Game Box Score**: View players box score from a selected game
- [ ] **Game Team Stats**: View team statistics from a selected game



[![Go Report Card](https://goreportcard.com/badge/github.com/connorvanderhook/nbaconsole?style=flat-square)](https://goreportcard.com/report/github.com/connorvanderhook/nbaconsole)