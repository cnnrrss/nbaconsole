package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cnnrrss/nbaconsole/api"
	"github.com/jroimartin/gocui"
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
	if len(nba.gamesList.Games) == 0 {
		return ""
	}

	nba.debuglog("idx " + string(idx))
	return nba.gamesList.Games[idx]
}

// HighlightedRowIndex returns the index of the highlighted row
func (nba *NBAConsole) HighlightedRowIndex() int {
	_, idx := nba.scoreboard.Origin()

	if idx < 0 {
		idx = 0
	}
	if idx >= len(nba.gamesList.Games) {
		idx = len(nba.gamesList.Games) - 1
	}
	return idx
}

func (nba *NBAConsole) getBoxScore() error {
	_, curH := nba.g.Size()

	params := genericParams(nba.date)
	if nba.selectedGame == "" {
		nba.debuglog("error nba selected game nil")
		panic(nba.selectedGame)
	}
	resp, err := api.GetDataGameBoxScore(params, nba.selectedGame)
	if err != nil {
		return fmt.Errorf("Error with boxscore request %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading request body %v", err)
	}

	boxScore := api.GameBoxScore{}
	if err := json.Unmarshal(body, &boxScore); err != nil {
		nba.debuglog(fmt.Sprintf("err unmarshalling %v\n", err.Error()))
	}

	nba.g.SetCurrentView("boxScore")

	nba.update(func() {
		nba.boxScore.Clear()
		fmt.Fprintln(nba.boxScore, boxScore.PointsLeaders())
		fmt.Fprintln(nba.boxScore, boxScore.AssistsLeaders())
		fmt.Fprintln(nba.boxScore, boxScore.ReboundsLeaders())
		nba.boxScore.SetCursor(0, curH-2)
		nba.boxScore.Highlight = true
		nba.boxScore.SelFgColor = gocui.ColorBlue
		nba.boxScore.SelBgColor = gocui.ColorDefault
	})
	return nil
}
