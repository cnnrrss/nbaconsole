package app

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

var (
	globalX0     int = 1
	scoreboardY0 int = 1
	footerHeight int = 2
)

var (
	headerLabel     string = "NBA Console"
	scoreboardLabel string = "scoreboard"
	boxScoreLabel   string = "boxscore"
	teamStatsLabel  string = "teamstats"
	footerLabel     string = "footer"
)

var (
	// MINWIDTH required to display the app
	MINWIDTH int = 50
	// MINHEIGHT required to display the app
	MINHEIGHT int = 30
)

func (nba *NBAConsole) layout(g *gocui.Gui) error {
	tw, th := g.Size()
	if nba.curW != tw || nba.curH != th {
		nba.refresh()
		nba.curW, nba.curH = min(tw, MINWIDTH), min(th, MINHEIGHT)
	}

	setMainView := func(g *gocui.Gui, fn func(g *gocui.Gui) error) {
		if err := fn(g); err != nil {
			nba.debuglog(fmt.Sprintf("error setting layout %v\n", err))
		}
	}
	setMainView(g, nba.setScoreboardView)
	setMainView(g, nba.setBoxScoreView)
	setMainView(g, nba.setFooterView)

	g.SetCurrentView(scoreboardLabel)

	return nil
}

func (nba *NBAConsole) setScoreboardView(g *gocui.Gui) error {
	if v, err := g.SetView(scoreboardLabel, 0, 7, nba.curW-1, nba.curH-(footerHeight*2)); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = fmt.Sprintf(" %-14s%s\n",
			headerLabel,
			toHumanDate(nba.date),
		)

		nba.scoreboard = v
		nba.scoreboard.Frame = false // maybe do this in the future
		highlightView(nba.scoreboard)

		if nba.gamesList == nil {
			scoreBoardBox := newBox(v, false)
			nba.gamesList = scoreBoardBox
		}

		if len(nba.gamesList.gameIDs) == 0 {
			go func() {
				nba.getScoreboard()
			}()
			nba.pollScoreboardData()
		}
	}
	return nil
}

func (nba *NBAConsole) setBoxScoreView(g *gocui.Gui) error {
	var err error
	if v, err := g.SetView(boxScoreLabel, 0 /** globalX0 */, 0 /** scoreboardY0 */, nba.curW-1, 6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		nba.boxScore = v
		nba.boxScore.Frame = false
		if nba.selectedGameID == "" {
			return fmt.Errorf("selectedGameID not yet set")
		}

		go func(err error) error {
			err = nba.getBoxScore()
			return err
		}(err)
	}

	return err
}

func (nba *NBAConsole) setFooterView(g *gocui.Gui) error {
	if v, err := g.SetView(footerLabel, globalX0, nba.curH-1-footerHeight, nba.curW-1, nba.curH-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nba.footerView = v
		nba.footerView.Frame = false
		go nba.updateFooter()
	}
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
