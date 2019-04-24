package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cnnrrss/nbaconsole/api"
	"github.com/jroimartin/gocui"
)

func (nba *NBAConsole) getBoxScore(gameID string) error {
	_, curH := nba.g.Size()

	params := genericParams(nba.date)
	resp, err := api.GetDataGameBoxScore(params, gameID)
	if err != nil {
		return fmt.Errorf("Error with boxscore request %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading request body %v", err)
	}

	boxScore := api.GameBoxScore{}
	json.Unmarshal(body, &boxScore)

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
