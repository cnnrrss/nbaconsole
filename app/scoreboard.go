package nbaconsole

import (
	"fmt"
	"io/ioutil"
	"sync"

	api "github.com/connorvanderhook/nbaconsole/api"
)

var scoreBoard, scoreBoardLock sync.Mutex

func (nba *NBAConsole) getScoreboard() error {
	scoreBoard.Lock()

	params := map[string]string{
		"DayOffset": "0",
		"LeagueID":  "00",
		"GameDate":  nba.date,
	}

	resp, err := api.ScoreboardV2(params)

	if err != nil {
		return fmt.Errorf("Error with request %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading request body %v", err)
	}

	nba.update(func() {
		nba.scoreboardView.Clear()
		fmt.Fprintln(nba.scoreboardView, "hello"+string(body))
	})

	scoreBoard.Unlock()

	return nil
}
