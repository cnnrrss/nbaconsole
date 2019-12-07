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

var scoreBoard sync.Mutex

var (
	scoreBoardHeaderFmt string = "%-11s%-14s%-13s%s\n"
	noGameScoresMsgFmt  string = "No hoops today, %s\n"
	gameScoreFmt        string = "%-9s%-3s%2s%4s%11s%15s\n"
)

// ToggleScoreboard toggles the scoreboard view and
// deletes any current view that may be displayed
func (nba *NBAConsole) ToggleScoreboard() error {
	// TODO: make better
	nba.g.DeleteView(teamStatsLabel)
	nba.g.DeleteView(boxScoreLabel)

	go nba.setScoreboardView(nba.g)

	return nil
}

// getScoreboard locks the current scoreBoard view and requests new data from
// the NBA API. If the data is received, the scoreBoard view is written to stdout
func (nba *NBAConsole) getScoreboard() error {
	// nba.scoreboard.Clear() don't clear its annoying
	nba.g.SetCurrentView(scoreboardLabel)

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
		nba.drawScoreboard(nba.scoreboard, sb, curW)
	})
	scoreBoard.Unlock()

	return nil
}

func (nba *NBAConsole) drawScoreboard(output io.Writer, sb api.DataScoreboard, width int /** TODO */) {
	nba.gamesList.Wipe()

	var blob strings.Builder
	headerString := fmt.Sprintf(scoreBoardHeaderFmt, "Home", "Score", "Away", "Status")
	blob.WriteString(headerString)
	blob.WriteString(
		fmt.Sprintf("%s\n",
			pad.AddString(len(headerString)-1, "-"),
		),
	)

	for _, gm := range sb.Games {
		blob.WriteString(formatGame(gm))
		nba.gamesList.games = append(nba.gamesList.games, gm.GameID)
	}

	if len(nba.gamesList.games) == 0 {
		blob.WriteString(
			fmt.Sprintf(noGameScoresMsgFmt,
				nba.message),
		)
	}

	fmt.Fprintf(output, blob.String())

	nba.scoreboard.SetOrigin(0, 0)
	nba.scoreboard.SetCursor(0, 2)

	highlightView(nba.scoreboard)
}

func formatGame(gm api.Game) string {
	var blob strings.Builder
	hScore, vScore := gm.Score()
	blob.WriteString(
		fmt.Sprintf(gameScoreFmt,
			gm.HTeam.TriCode,
			hScore,
			"-",
			vScore,
			gm.VTeam.TriCode,
			gm.Status(),
		),
	)

	if gm.Playoffs.RoundNum != "" {
		// TODO: design new fmt
		blob.WriteString(pad.Left(gm.Playoffs.SeriesSummaryText, 6, " "))
	}

	return blob.String()
}
