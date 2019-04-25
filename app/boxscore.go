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
	if nba.selectedGame == &selectedGame {
		nba.selectedGame = nil
	} else {
		nba.selectedGame = &selectedGame
	}

	go nba.setBoxScoreView(nba.g, nba.selectedGame.ID)

	return nil
}

// HighlightedRowCoin returns the coin at the index of the highlighted row
func (nba *NBAConsole) SelectedGame() GameScore {
	idx := nba.HighlightedRowIndex()
	if len(nba.gamesList.Items) == 0 {
		return GameScore{}
	}

	return *nba.gamesList.Items[idx]
}

// HighlightedRowIndex returns the index of the highlighted row
func (nba *NBAConsole) HighlightedRowIndex() int {
	_, y := nba.scoreboard.Origin()
	_, cy := nba.scoreboard.Cursor()
	idx := y + cy
	if idx < 0 {
		idx = 0
	}
	if idx >= len(nba.gamesList.Items) {
		idx = len(nba.gamesList.Items) - 1
	}
	return idx
}

func (nba *NBAConsole) getBoxScore() error {
	_, curH := nba.g.Size()

	params := genericParams(nba.date)
	if nba.selectedGame == nil {
		nba.debuglog("error nba selected game nil")
		panic(nba.selectedGame)
	}

	resp, err := api.GetDataGameBoxScore(params, nba.selectedGame.ID)
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
		nba.scoreboard.Clear()
		fmt.Fprintln(nba.boxScore, boxScore.PointsLeaders())
		nba.boxScore.SetCursor(0, curH-2)
		nba.boxScore.Highlight = true
		nba.boxScore.SelFgColor = gocui.ColorBlue
		nba.boxScore.SelBgColor = gocui.ColorDefault
	})
	return nil
}
