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

func (nba *NBAConsole) toggleScoreboard() error {
	var err error
	if nba.scoreboard != nil && nba.boxScore != nil {
		nba.teamStats.Clear() // TODO: make better

		_, err = nba.g.SetViewOnTop(scoreboardLabel) // todo: make better
		_, err = nba.g.SetViewOnTop(boxScoreLabel)
		return err
	}

	go nba.getScoreboard()

	return nil
}

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
	if err := json.Unmarshal(body, &sb); err != nil {
		return fmt.Errorf("Could not unmarshal response body to data scoreboard struct %v", err)
	}

	scoreBoard.Lock()
	nba.update(func() {
		curW, _ := nba.g.Size()
		nba.drawScoreboard(nba.scoreboard, sb, curW)
	})
	scoreBoard.Unlock()

	return nil
}

func (nba *NBAConsole) drawScoreboard(output io.Writer, sb api.DataScoreboard, width int /** TODO */) {
	nba.gamesList.wipe() // don't redraw the games

	var scoreboardBlob strings.Builder
	headerString := fmt.Sprintf(scoreBoardHeaderFmt, "Home", "Score", "Away", "Status")
	scoreboardBlob.WriteString(headerString)
	scoreboardBlob.WriteString(
		fmt.Sprintf("%s\n",
			pad.AddString(len(headerString)-1, "-"),
		),
	)

	for _, gm := range sb.Games {
		scoreboardBlob.WriteString(formatGames(gm))
		nba.gamesList.gameIDs = append(nba.gamesList.gameIDs, gm.GameID)
	}

	if len(nba.gamesList.gameIDs) > 0 {
		nba.selectedGameID = nba.gamesList.gameIDs[0]
	}

	// TODO: this call probs shouldn't be here
	// go nba.getBoxScore() choose not to pull
	// in a goroutine because it renders cleaner
	if err := nba.getBoxScore(); err != nil {
		nba.debuglog(fmt.Sprintf("Could not get boxScore. err: %v\n", err))
	}

	if len(nba.gamesList.gameIDs) == 0 {
		scoreboardBlob.WriteString(
			fmt.Sprintf(noGameScoresMsgFmt,
				nba.message),
		)
	}

	if _, err := fmt.Fprintf(output, scoreboardBlob.String()); err != nil {
		nba.debuglog(fmt.Sprintf("Could not print scoreboard. err: %v\n", err))
	}

	if err := nba.g.CurrentView().SetOrigin(0, 0); err != nil {
		nba.debuglog(fmt.Sprintf("Could not set origin on scoreboard: %v\n", err))
	}

	if err := nba.g.CurrentView().SetCursor(0, 2); err != nil {
		nba.debuglog(fmt.Sprintf("Could not set cursor on scoreboard: %v\n", err))
	}

	highlightView(nba.scoreboard)
}

func formatGames(gm api.Game) string {
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
