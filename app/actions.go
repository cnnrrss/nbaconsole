package app

import (
	"github.com/jroimartin/gocui"
	log "github.com/sirupsen/logrus"
)

// refresh reads from the rateLimiter channel
// and write to the forceRefresh Channel
func (nba *NBAConsole) refresh() error {
	go func() {
		<-nba.rateLimiter
		nba.forceRefresh <- true
	}()
	return nil
}

// update passes in a function to update an NBAConsole view
func (nba *NBAConsole) update(f func()) {
	if nba.g == nil {
		log.Fatal("gocui is not initialized")
	}

	nba.g.Update(func(g *gocui.Gui) error {
		f()
		return nil
	})
}

// pollScoreBoardData is a go routine that listens on 2 channels:
// 1) the forceRefresh directive, 2) the refreshTicker internal Timer
// when a message is read from these channels a refresh is triggered
func (nba *NBAConsole) pollScoreboardData() {
	go func() {
		for {
			select {
			case <-nba.forceRefresh:
				nba.refreshScoreboardView()
				go nba.updateFooter("")
			case <-nba.refreshTicker.C:
				nba.refreshScoreboardView()
				go nba.updateFooter("")
			}
		}
	}()
}

// refreshScoreBoardView is a go routine to clear the current
// scoreboard and make a get request to refresh the latest data
func (nba *NBAConsole) refreshScoreboardView() error {
	go func() {
		nba.getScoreboard()
	}()
	return nil
}
