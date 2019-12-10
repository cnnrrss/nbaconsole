package app

import (
	"github.com/jroimartin/gocui"
)

func (nba *NBAConsole) getSelectedGameID() string {
	idx := nba.highlightedRowIndex()
	if len(nba.gamesList.gameIDs) == 0 {
		nba.debuglog("gameslist is empty, cant find idx " + string(idx))
		return ""
	}
	return nba.gamesList.gameIDs[idx]
}

func (nba *NBAConsole) highlightedRowIndex() int {
	_, oy := nba.scoreboard.Origin()
	_, y := nba.scoreboard.Cursor()

	// Skip 2 static lines in the scoreboard view
	idx := y - 2 - oy

	if idx < 0 {
		idx = 0
	}

	if idx >= len(nba.gamesList.gameIDs) {
		idx = len(nba.gamesList.gameIDs) - 1
	}
	return idx
}

func highlightView(v *gocui.View) {
	v.Highlight = true
	v.SelFgColor = gocui.ColorBlue
	v.SelBgColor = gocui.ColorDefault
}
