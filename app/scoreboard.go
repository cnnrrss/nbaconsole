package nbaconsole

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	api "github.com/connorvanderhook/nbaconsole/api"
)

var scoreBoard, scoreBoardLock sync.Mutex

func (nba *NBAConsole) getScoreboard(params map[string]string) error {
	scoreBoard.Lock()

	resp, err := api.GetDataScoreBoard(params)
	if err != nil {
		fmt.Fprintf(nba.scoreboard, "Error happened on get scoreboard")
		return fmt.Errorf("Error with request %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(nba.scoreboard, "Error happened on get scoreboard")
		return fmt.Errorf("Error reading request body %v", err)
	}

	sb := api.DataScoreboard{}
	json.Unmarshal(body, &sb)
	curW, _ := nba.g.Size()

	nba.update(func() {
		nba.scoreboard.Clear()
		nba.setGames(sb)
		nba.OuputScoreBoard(curW)
	})

	scoreBoard.Unlock()

	return nil
}

// OuputScoreBoard prints the current games to the scoreboard view
func (nba *NBAConsole) OuputScoreBoard(width int) {
	if len(nba.gamesList.games) > 0 {
		fmt.Fprintln(nba.scoreboard, formatScoreBoardHeader(width-2))
		fmt.Fprintln(nba.scoreboard, boxSeparator(width-2))
		for i, game := range nba.gamesList.games {
			item := fmt.Sprintf("%s\n", game)
			fmt.Fprintf(nba.scoreboard, string(i)+item)
			fmt.Fprintln(nba.scoreboard, boxSeparator(width-2))
		}
		return
	}
	fmt.Fprintf(nba.scoreboard, "Sorry no hoops today.. %v\n", nba.date)
	return
}

// genericParams returns basic parameters to include in an API request
func genericParams(date string) map[string]string {
	params := map[string]string{
		"DayOffset": "0",
		"LeagueID":  "00",
		"gameDate":  date,
	}
	return params
}

func formatScoreBoardHeader(width int) string {
	var header string
	for _, h := range []string{"Away", "Home", "Score", "Status"} {
		header += PadCenter(h, width/4)
	}
	return header
}

func (nba *NBAConsole) setGames(sb api.DataScoreboard) error {
	data := make([]interface{}, len(sb.Games))
	curX, _ := nba.g.Size()
	for i, gm := range sb.Games {
		var blob string
		hScore, vScore := gm.Score()
		blob += PadCenter(gm.VTeam.TriCode, (curX/4)-3)
		blob += PadCenter(gm.HTeam.TriCode, (curX/4)-3)
		scoreOffset := len(hScore) + len(vScore) + 1
		if len(vScore) == 2 {
			vScore = "_" + vScore
		}
		if len(hScore) == 2 {
			hScore = "_" + hScore
		}
		blob += PadCenter(fmt.Sprintf("%s - %s", vScore, hScore), (curX/4)+scoreOffset)
		blob += PadCenter(gm.Status(), (curX/4)-scoreOffset)

		if gm.Playoffs.RoundNum != "" {
			blob += fmt.Sprintf("%s", PadLeft(gm.Playoffs.SeriesSummaryText, 6))
		}
		data[i] = blob
	}

	nba.gamesList.games = data
	nba.gamesList.curIndex = 0
	return nil
}

func boxSeparator(w int) string {
	return fmt.Sprintf(Pad("-", w-1))
}
