package app

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/cnnrrss/nbaconsole/api"
)

// ToggleGameBoxScore toggles between the global scoreboard and the game box score
func (nba *NBAConsole) ToggleGameBoxScore() error {
	selectedGame := nba.SelectedGame()
	if nba.selectedGame != selectedGame {
		nba.selectedGame = selectedGame
	}

	go nba.setBoxScoreView(nba.g, nba.selectedGame)

	return nil // TODO: handle errors gracefully
}

// SelectedGame returns the game of the highlighted row
func (nba *NBAConsole) SelectedGame() string {
	idx := nba.HighlightedRowIndex()
	if len(nba.gamesList.games) == 0 {
		return ""
	}
	return nba.gamesList.games[idx]
}

// HighlightedRowIndex returns the index of the highlighted row
func (nba *NBAConsole) HighlightedRowIndex() int {
	_, oy := nba.scoreboard.Origin()
	_, y := nba.scoreboard.Cursor()
	// Skip 2 static lines in the scoreboard view
	idx := y - 2 - oy

	if idx < 0 {
		idx = 0
	}

	if idx >= len(nba.gamesList.games) {
		idx = len(nba.gamesList.games) - 1
	}
	return idx
}

func (nba *NBAConsole) getBoxScore() error {
	nba.g.SetCurrentView(boxScoreLabel)

	params := genericParams(nba.date)
	if nba.selectedGame == "" {
		nba.debuglog("error nba selected game nil")
		panic(nba.selectedGame)
	}

	gameBoxScore := api.GameBoxScore{}

	// if final game not cached. // TODO: should cache all games.
	if nba.selectedGameScore == nil ||
		nba.selectedGameScore.ID != nba.selectedGame ||
		nba.selectedGameScore.SportsContent.Game.PeriodTime.GameStatus == "3" {

		resp, err := api.GetDataGameBoxScore(params, nba.selectedGame)
		if err != nil {
			return fmt.Errorf("Error with boxscore request %v", err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Error reading request body %v", err)
		}

		if err := json.Unmarshal(body, &gameBoxScore); err != nil {
			nba.debuglog(fmt.Sprintf("err unmarshalling %v\n", err.Error()))
		}

		nba.selectedGameScore = &GameScore{
			gameBoxScore,
			nba.selectedGame,
		}
	}

	nba.update(func() {
		nba.boxScore.Clear()
		nba.drawBoxScore(nba.boxScore, nba.selectedGameScore, 0)
	})
	return nil
}

func (nba *NBAConsole) drawBoxScore(output io.Writer, bs *GameScore, width int) {
	var str strings.Builder
	str.WriteString(fmt.Sprintf("%-25s%-4s%-4s%-4s%-5s%s\n", "Team", "1", "2", "3", "4", "T"))
	hLine, hTotal := HomeLineScores(bs)
	str.WriteString(
		fmt.Sprintf("%-23s%-5s%4d\n",
			bs.GameBoxScore.SportsContent.Game.Home.City+" "+
				bs.GameBoxScore.SportsContent.Game.Home.Nickname,
			hLine,
			hTotal,
		),
	)
	vLine, vTotal := VisitorLineScores(bs)
	str.WriteString(
		fmt.Sprintf("%-23s%-5s%4d\n",
			bs.GameBoxScore.SportsContent.Game.Visitor.City+" "+
				bs.GameBoxScore.SportsContent.Game.Visitor.Nickname,
			vLine,
			vTotal,
		),
	)

	fmt.Fprintln(output, str.String())
	fmt.Fprintln(output, bs.BoxScoreLeaders())
	highlightView(nba.scoreboard)
}

// HomeLineScores ...
func HomeLineScores(bs *GameScore) (string, int) {
	var lineScore strings.Builder
	var total int
	curPeriod := len(bs.GameBoxScore.SportsContent.Game.Home.Linescores.Period)

	for i := 0; i < curPeriod || i <= 3; i++ {
		if i >= curPeriod {
			lineScore.WriteString("  - ")
		} else {
			q, _ := strconv.Atoi(bs.GameBoxScore.SportsContent.Game.Home.Linescores.Period[i].Score)
			total += q
			lineScore.WriteString(
				fmt.Sprintf("%2d  ", q),
			)
		}
	}
	return lineScore.String(), total
}

// VisitorLineScores ...
func VisitorLineScores(bs *GameScore) (string, int) {
	var lineScore strings.Builder
	var total int
	curPeriod := len(bs.GameBoxScore.SportsContent.Game.Visitor.Linescores.Period)

	for i := 0; i < curPeriod || i <= 3; i++ {
		if i >= curPeriod {
			lineScore.WriteString("  - ")
		} else {
			q, _ := strconv.Atoi(bs.GameBoxScore.SportsContent.Game.Visitor.Linescores.Period[i].Score)
			total += q
			lineScore.WriteString(
				fmt.Sprintf("%2d  ", q),
			)
		}
	}
	return lineScore.String(), total
}
