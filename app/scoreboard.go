package nbaconsole

import (
	"fmt"
	"log"
	"sync"
)

var scoreBoard, scoreBoardLock sync.Mutex
var options = "" // no options for now

func (nba *NBAConsole) getScoreboard() error {
	scoreBoard.Lock()

	results, err := nba.client.TodaysGames(options)
	if err != nil {
		log.Fatalf("couldn't fetch data")
	}

	content := fmt.Sprintf("%s", results)

	nba.update(func() {
		nba.scoreboardView.Clear()
		fmt.Fprintln(nba.scoreboardView, content)
	})

	scoreBoard.Unlock()

	return nil
}
