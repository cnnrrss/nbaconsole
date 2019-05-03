package app

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

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

// HighlightedRowCoin returns the coin at the index of the highlighted row
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
	fmt.Fprintln(output, bs.PointsLeaders())
	fmt.Fprintln(output, bs.AssistsLeaders())
	fmt.Fprintln(output, bs.ReboundsLeaders())
	highlightView(nba.scoreboard)
}
