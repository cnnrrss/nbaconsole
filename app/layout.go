package nbaconsole

import (
	"fmt"

	"github.com/jroimartin/gocui"
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

	g.SetCurrentView("scoreboard")

	return nil
}

func (nba *NBAConsole) setScoreboardLayout(g *gocui.Gui, params map[string]string) error {
	if v, err := g.SetView("scoreboard", 1, 3, nba.curW-1, nba.curH-1); err != nil {
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

func (nba *NBAConsole) setHeaderLayout(g *gocui.Gui) error {
	if v, err := g.SetView("welcome", 1, 0, nba.curW-1, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		fmt.Fprintf(v, "%s\n", PadCenter("Welcome to NBA Console", nba.curW/2))
	}
	return nil
}

func (nba *NBAConsole) setFooterLayout(g *gocui.Gui) error {
	if v, err := g.SetView("footerview", 1, 3, nba.curW-1, 4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nba.footerView = v
		nba.footerView.Frame = true
		go nba.updateFooter("")
	}
	return nil
}
