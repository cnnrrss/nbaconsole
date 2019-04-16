package nbaconsole

import (
	"github.com/jroimartin/gocui"
)

func (nba *NBAConsole) scoreBoardLayout(g *gocui.Gui) error {
	if v, err := g.SetView("scoreboard", 0, 3, 40, 15); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nba.scoreboardView = v
		nba.scoreboardView.Frame = true
		nba.scoreboardView.Editable = false
		nba.scoreboardView.Wrap = true
		nba.scoreboardView.Highlight = true
		nba.scoreboardView.BgColor = gocui.ColorBlack
		nba.scoreboardView.FgColor = gocui.ColorMagenta

		go func() {
			// Make call to NBA API in goroutine
			nba.getScoreboard()
		}()
	}
	return nil
}
