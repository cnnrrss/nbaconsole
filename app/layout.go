package nbaconsole

import (
	"github.com/jroimartin/gocui"
)

const scoreBoardLayoutName = "scoreboard"

func (nba *NBAConsole) scoreBoardLayout(g *gocui.Gui) error {
	maxX, _ := nba.size()
	topOffset := 0

	if v, err := g.SetView(scoreBoardLayoutName, 0, topOffset, maxX, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nba.scoreboardView = v
		nba.scoreboardView.Frame = false
		nba.scoreboardView.BgColor = gocui.ColorBlack
		nba.scoreboardView.FgColor = gocui.ColorWhite

		go func() {
			// Make call to NBA API in goroutine
			nba.getScoreboard()
		}()
	}
	return nil
}
