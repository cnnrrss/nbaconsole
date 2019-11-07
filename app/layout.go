package app

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

var (
	globalX0 = 1

	headerLabel = "NBA Console"

	scoreboardLabel = "scoreboard"
	scoreboardY0    = 1

	boxScoreLabel  = "boxscore"
	teamStatsLabel = "teamstats"

	footerLabel  = "footer"
	footerHeight = 2
)

func (nba *NBAConsole) layout(g *gocui.Gui) error {
	tw, th := g.Size()
	if nba.curW != tw || nba.curH != th {
		nba.refresh()
		nba.curW, nba.curH = tw, th
	}

	setMainView := func(g *gocui.Gui, fn func(g *gocui.Gui) error) {
		if err := fn(g); err != nil {
			log.Printf("error setting layout %v\n", err)
		}
	}

	setMainView(g, nba.setScoreboardView)
	setMainView(g, nba.setFooterView)

	g.SetCurrentView(scoreboardLabel)

	return nil
}

func (nba *NBAConsole) setScoreboardView(g *gocui.Gui) error {
	if v, err := g.SetView(scoreboardLabel, 0, 0, nba.curW-1, nba.curH-footerHeight-footerHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = fmt.Sprintf(" %-14s%s\n",
			headerLabel,
			toHumanDate(nba.date),
		)

		nba.scoreboard = v
		highlightView(nba.scoreboard)
		scoreBoardBox := NewBox(v, false)
		nba.gamesList = scoreBoardBox

		go func() {
			nba.getScoreboard()
		}()
		nba.pollScoreboardData()
	}
	return nil
}

func (nba *NBAConsole) setBoxScoreView(g *gocui.Gui, gameID string) error {
	if v, err := g.SetView(boxScoreLabel, globalX0, scoreboardY0, nba.curW-1, nba.curH-footerHeight-footerHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		nba.boxScore = v

		go func() {
			nba.getBoxScore()
		}()
	}
	return nil
}

func (nba *NBAConsole) setFooterView(g *gocui.Gui) error {
	if v, err := g.SetView(footerLabel, globalX0, nba.curH-1-footerHeight, nba.curW-1, nba.curH-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nba.footerView = v
		nba.footerView.Frame = true
		go nba.updateFooter("")
	}
	return nil
}
