package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/cnnrrss/nbaconsole/api"
	"github.com/cnnrrss/nbaconsole/common/pad"
	"github.com/jroimartin/gocui"
)

var (
	scoreBoard, scoreBoardLock sync.Mutex
	scoreBoardHeaders          = []string{"Home", "Away", "Score", "Status"}
)

// getScoreboard locks the current scoreBoard view and requests new data from
// the NBA API. If the data is received, the scoreBoard view is written to stdout
func (nba *NBAConsole) getScoreboard() error {
	scoreBoard.Lock()

	curW, _ := nba.g.Size()
	params := genericParams(nba.date)
	resp, err := api.GetDataScoreBoard(params)
	if err != nil {
		return fmt.Errorf("Error with request %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer body.Close()
	if err != nil {
		return fmt.Errorf("Error reading request body %v", err)
	}

	sb := api.DataScoreboard{}
	json.Unmarshal(body, &sb)
	nba.update(func() {
		nba.scoreboard.Clear()
		nba.setGames(sb) // TODO: looping twice
		nba.OuputScoreBoard(curW)
		_, y := nba.scoreboard.Cursor()
		nba.scoreboard.SetCursor(0, y+2)
		nba.scoreboard.Highlight = true
		nba.scoreboard.SelFgColor = gocui.ColorBlue
		nba.scoreboard.SelBgColor = gocui.ColorDefault
	})

	scoreBoard.Unlock()

	return nil
}

// OuputScoreBoard prints the current games to the scoreboard view
func (nba *NBAConsole) OuputScoreBoard(width int) {
	if len(nba.gamesList.Items) > 0 {
		fmt.Fprintln(nba.scoreboard, formatScoreBoardHeader(width-2))
		fmt.Fprintln(nba.scoreboard, pad.Left(fmt.Sprint("-"), nba.curW-1, "-"))
		for _, g := range nba.gamesList.Items {
			fmt.Fprintln(nba.scoreboard, g)
		}
		return
	}
	fmt.Fprintf(nba.scoreboard, "No hoops today, %s\n", nba.message)
	return
}

func formatScoreBoardHeader(width int) string {
	var header string
	for i, h := range scoreBoardHeaders {
		switch {
		case i < len(scoreBoardHeaders)-1:
			header += pad.Left(h, 6+(1*i), " ")
		default:
			header += pad.Left(h, 12, " ")
		}
	}
	return header
}

func (nba *NBAConsole) setGames(sb api.DataScoreboard) error {
	data := make([]interface{}, len(sb.Games))
	for i, gm := range sb.Games {
		var blob string
		hScore, vScore := gm.Score()
		blob += pad.Left(gm.VTeam.TriCode, 5, " ") // TODO: fix hardcoding
		blob += pad.Left(gm.HTeam.TriCode, 7, " ")
		blob += pad.Left(fmt.Sprintf("%s - %s", vScore, hScore), len(blob), " ")
		blob += pad.Left(gm.Status(), 8, " ")
		if gm.Playoffs.RoundNum != "" {
			blob += fmt.Sprintf("%s", pad.Left(gm.Playoffs.SeriesSummaryText, 6, " "))
		}
		data[i] = blob
	}
	nba.gamesList.Items = data
	nba.gamesList.CurrentIndex = 0
	return nil
}
