package nbaconsole

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	api "github.com/connorvanderhook/nbaconsole/api"
)

var scoreBoard, scoreBoardLock sync.Mutex

func (nba *NBAConsole) getScoreboard() error {
	scoreBoard.Lock()

	// TODO: pass in params
	params := map[string]string{
		"DayOffset": "0",
		"LeagueID":  "00",
		"gameDate":  "20190318",
	}

	resp, err := api.GetDataScoreBoard(params)
	if err != nil {
		fmt.Fprintf(nba.scoreboardView, "Error happened on get scoreboard")
		return fmt.Errorf("Error with request %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(nba.scoreboardView, "Error happened on get scoreboard")
		return fmt.Errorf("Error reading request body %v", err)
	}

	sb := api.DataScoreboard{}
	json.Unmarshal(body, &sb)

	nba.update(func() {
		nba.scoreboardView.Clear()
		fmt.Fprintf(nba.scoreboardView, "Games %s\n", nba.date)
		fmt.Fprintln(nba.scoreboardView, formatGames(sb))
	})

	scoreBoard.Unlock()

	return nil
}

func formatGames(sb api.DataScoreboard) string {
	var resp string
	for _, game := range sb.Games {
		resp += fmt.Sprintf("----------------------------------\n")
		if game.Playoffs.RoundNum != "" {
			// if round num = 3 "Conference Finals"
			resp += fmt.Sprintf("%s\n", game.Playoffs.SeriesSummaryText)
		}
		// TODO: handle overtimes
		resp += fmt.Sprintf("Teams | Q1 | Q2 | Q3 | Q4 | Total\n")
		resp += fmt.Sprintf("%-5s | ", game.HTeam.TriCode)
		for _, q := range game.HTeam.Linescore {
			resp += fmt.Sprintf(" %3v|", q.Score)
		}
		resp += fmt.Sprintf(" %s \n", game.HTeam.Score)
		resp += fmt.Sprintf("%-5s | ", game.VTeam.TriCode)
		for _, q := range game.VTeam.Linescore {
			resp += fmt.Sprintf(" %3v|", q.Score)
		}
		resp += fmt.Sprintf(" %s \n", game.VTeam.Score)
	}
	return resp
}
