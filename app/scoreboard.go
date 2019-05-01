package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/cnnrrss/nbaconsole/api"
	"github.com/cnnrrss/nbaconsole/common/pad"
	"github.com/jroimartin/gocui"
)

var (
	scoreBoard        sync.Mutex
	scoreBoardHeaders = []string{"Home", "Away", "Score", "Status"}
)

// getScoreboard locks the current scoreBoard view and requests new data from
// the NBA API. If the data is received, the scoreBoard view is written to stdout
func (nba *NBAConsole) getScoreboard() error {

	params := genericParams(nba.date)
	resp, err := api.GetDataScoreBoard(params)
	if err != nil {
		return fmt.Errorf("Error with request %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading request body %v", err)
	}

	sb := api.DataScoreboard{}
	json.Unmarshal(body, &sb)

	scoreBoard.Lock()
	nba.update(func() {
		curW, _ := nba.g.Size()
		nba.scoreboard.Clear()
		nba.drawScoreboard(sb, curW)
		_, y := nba.scoreboard.Cursor()
		nba.scoreboard.SetCursor(0, y+2)
		nba.scoreboard.SelFgColor = gocui.ColorBlue
		nba.scoreboard.SelBgColor = gocui.ColorDefault
	})
	scoreBoard.Unlock()

	return nil
}

// DrawScoreBoard prints the current games to the scoreboard view
func (nba *NBAConsole) DrawScoreBoard(width int) {
	if len(nba.gamesList.Games) > 0 {
		fmt.Fprintln(nba.scoreboard, formatScoreBoardHeader(width-2))
		fmt.Fprintln(nba.scoreboard, pad.Left(fmt.Sprint("-"), nba.curW-1, "-"))
		for _, g := range nba.gamesList.Games {
			fmt.Fprintln(nba.scoreboard, g)
		}
		return
	}
	fmt.Fprintf(nba.scoreboard, "No hoops today, %s\n", nba.message)
}

func formatScoreBoardHeader(width int) string {
	var str strings.Builder
	for i, h := range scoreBoardHeaders {
		switch {
		case i < len(scoreBoardHeaders)-1:
			str.WriteString(pad.Left(h, 6+(1*i), " "))
		default:
			str.WriteString(pad.Left(h, 12, " "))
		}
	}
	return str.String()
}

func (nba *NBAConsole) drawScoreboard(sb api.DataScoreboard, width int) error {
	nba.gamesList.Wipe()
	fmt.Fprintln(nba.scoreboard, formatScoreBoardHeader(width))
	for _, gm := range sb.Games {
		var blob strings.Builder
		hScore, vScore := gm.Score()
		blob.WriteString(pad.Left(gm.VTeam.TriCode, 5, " "))
		blob.WriteString(pad.Left(gm.HTeam.TriCode, 7, " "))
		blob.WriteString(pad.Left(fmt.Sprintf("%s - %s", vScore, hScore), blob.Len(), " "))
		blob.WriteString(pad.Left(gm.Status(), 8, " "))
		if gm.Playoffs.RoundNum != "" {
			blob.WriteString(pad.Left(gm.Playoffs.SeriesSummaryText, 6, " "))
		}
		fmt.Fprintln(nba.scoreboard, blob.String())
		nba.gamesList.Games = append(nba.gamesList.Games, gm.GameID)
	}
	nba.gamesList.CurrentIndex = 0
	return nil
}
