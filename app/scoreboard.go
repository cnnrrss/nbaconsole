package app

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/cnnrrss/nbaconsole/api"
	"github.com/cnnrrss/nbaconsole/common/pad"
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
	if err != nil {
		return fmt.Errorf("Error reading request body %v", err)
	}

	sb := api.DataScoreboard{}
	json.Unmarshal(body, &sb)
	nba.update(func() {
		nba.scoreboard.Clear()
		nba.setGames(sb) // TODO: looping twice
		nba.DrawScoreBoard(nba.scoreboard, curW)
		_, y := nba.scoreboard.Cursor()
		nba.scoreboard.SetCursor(0, y+2)
		nba.scoreboard.Highlight = true
	})

	scoreBoard.Unlock()

	return nil
}

// DrawScoreBoard prints the current games to the scoreboard view
func (nba *NBAConsole) DrawScoreBoard(output io.Writer, width int) {
	if len(nba.gamesList.Items) > 0 {
		fmt.Fprintln(output, formatScoreBoardHeader(width-2))
		fmt.Fprintln(output, pad.Left(fmt.Sprint("-"), nba.curW-1, "-"))
		for _, g := range nba.gamesList.Items {
			fmt.Fprintln(output, g.Msg)
		}
		return
	}
	fmt.Fprintf(output, "No hoops today, %s\n", nba.message)
}

func formatScoreBoardHeader(width int) string {
	var str strings.Builder
	for i, h := range scoreBoardHeaders {
		switch {
		case i < len(scoreBoardHeaders)-1:
			str.WriteString(pad.Left(h, 6+(1*i), " "))
		default:
			str.WriteString(pad.Left(h, 17, " "))
		}
	}
	return str.String()
}

func (nba *NBAConsole) setGames(sb api.DataScoreboard) {
	for _, gm := range sb.Games {
		var blob strings.Builder
		hScore, vScore := gm.Score()
		blob.WriteString(pad.Left(gm.VTeam.TriCode, 5, " "))
		blob.WriteString(pad.Left(gm.HTeam.TriCode, 7, " "))
		blob.WriteString(pad.Left(fmt.Sprintf("%s - %s", vScore, hScore), blob.Len(), " "))
		blob.WriteString(pad.Left(gm.Status(), 14, " "))
		if gm.Playoffs.RoundNum != "" {
			blob.WriteString(fmt.Sprintf("%s", pad.Left(gm.Playoffs.SeriesSummaryText, 6, " ")))
		}
		nba.gamesList.Items = append(nba.gamesList.Items, &GameScore{Msg: blob.String(), ID: gm.GameID})
	}
	nba.gamesList.CurrentIndex = 0
}
