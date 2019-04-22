package app

import (
	"fmt"

	"github.com/cnnrrss/nbaconsole/common/pad"
	"github.com/jroimartin/gocui"
)

var (
	globalLayout = "Welcome to NBA Console"

	globalX0 = 1

	headerLabel  = "welcome"
	headerY0     = 0
	headerHeight = 2
	headerY1     = headerY0 + headerHeight

	scoreboardLabel = "scoreboard"
	scoreboardY0    = headerY1 + 1

	footerLabel  = "footer"
	footerHeight = 2
)

//
func (nba *NBAConsole) layout(g *gocui.Gui) error {
	var err error
	// terminal width and height
	tw, th := g.Size()

	// set current app width and height
	if nba.curW != tw || nba.curH != th {
		nba.refresh()
		nba.curW = tw
		nba.curH = th
	}

	err = nba.setHeaderView(nba.g)
	if err != nil {
		return err
	}

	err = nba.setScoreboardView(nba.g)
	if err != nil {
		return err
	}

	err = nba.setFooterLayout(nba.g)
	if err != nil {
		return err
	}

	g.SetCurrentView(scoreboardLabel)

	return nil
}

func (nba *NBAConsole) setHeaderView(g *gocui.Gui) error {
	if v, err := g.SetView(headerLabel, globalX0, headerY0, nba.curW-1, headerY1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		dateString := toHumanDateTime(nba.date)
		fmt.Fprintf(v, " %s %s\n", globalLayout, pad.Left(dateString, len(globalLayout)+6, " ")) // TODO: no hardcode
	}
	return nil
}

func (nba *NBAConsole) setScoreboardView(g *gocui.Gui) error {
	if v, err := g.SetView(scoreboardLabel, globalX0, scoreboardY0, nba.curW-1, nba.curH-footerHeight-footerHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		nba.scoreboard = v
		nba.scoreboard.FgColor = gocui.ColorMagenta
		scoreBoardBox := NewBox(v, false)
		nba.gamesList = scoreBoardBox

		go func() {
			nba.getScoreboard()
		}()
		nba.pollScoreboardData()
	}
	return nil
}

func (nba *NBAConsole) setFooterLayout(g *gocui.Gui) error {
	if v, err := g.SetView("footerview", globalX0, nba.curH-1-footerHeight, nba.curW-1, nba.curH-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nba.footerView = v
		nba.footerView.Frame = true
		go nba.updateFooter("")
	}
	return nil
}
