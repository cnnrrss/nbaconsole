package app

import (
	"fmt"

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

func (nba *NBAConsole) layout(g *gocui.Gui) error {
	var err error
	tw, th := g.Size()
	params := genericParams(nba.date)

	if nba.curW != tw || nba.curH != th {
		nba.refreshAll(params)
		nba.curW = tw
		nba.curH = th
	}

	err = nba.setHeaderLayout(nba.g)
	if err != nil {
		return err
	}

	err = nba.setScoreboardLayout(nba.g, params)
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

func (nba *NBAConsole) setHeaderLayout(g *gocui.Gui) error {
	if v, err := g.SetView(headerLabel, globalX0, headerY0, nba.curW-1, headerY1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		fmt.Fprintf(v, "%s\n", PadCenter(globalLayout, nba.curW/2))
	}
	return nil
}

func (nba *NBAConsole) setScoreboardLayout(g *gocui.Gui, params map[string]string) error {
	if v, err := g.SetView(scoreboardLabel, globalX0, scoreboardY0, nba.curW-1, nba.curH-footerHeight-footerHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		nba.scoreboard = v
		scoreBoardBox := NewBox(v, false)
		nba.gamesList = scoreBoardBox

		go func() {
			// TODO: make this output to channel
			nba.getScoreboard(params)
		}()
		nba.pollScoreboardData(params)
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
