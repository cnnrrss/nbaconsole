package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cnnrrss/nbaconsole/api"
	"github.com/jroimartin/gocui"
)

func (nba *NBAConsole) getBoxScore(gameID string) error {
	curW, _ := nba.g.Size()

	params := genericParams(nba.date)
	resp, err := api.GetDataGameBoxScore(params, gameID)
	if err != nil {
		return fmt.Errorf("Error with boxscore request %v", err)
	}

	body, err := ioutil.ReadAll(resp.body)
	defer body.Close()
	if err != nil {
		return fmt.Errorf("Error reading request body %v", err)
	}

	boxScore := api.GameBoxScore{}
	json.Unmarshal(body, &boxScore)

	nba.update(func() {
		nba.scoreboard.Clear()
		hm, aw := boxScore.Teams()
		boxScore.PointsLeaders(hm, aw)
		nba.scoreboard.SetCursor(0, y+2)
		nba.scoreboard.Highlight = true
		nba.scoreboard.SelFgColor = gocui.ColorBlue
		nba.scoreboard.SelBgColor = gocui.ColorDefault
	})
}
