package app

import (
	"github.com/jroimartin/gocui"
)

func (nba *NBAConsole) keybindings(g *gocui.Gui) error {
	var err error

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlR, gocui.ModNone, nba.keyfn(nba.toggleScoreboard)); err != nil {
		return err
	}

	if err = g.SetKeybinding(scoreboardLabel, gocui.KeyArrowUp, gocui.ModNone, nba.keyfn(nba.moveUp)); err != nil {
		return err
	}

	if err = g.SetKeybinding(scoreboardLabel, gocui.KeyArrowDown, gocui.ModNone, nba.keyfn(nba.moveDown)); err != nil {
		return err
	}

	//if err := g.SetKeybinding(scoreboardLabel, gocui.KeyEnter, gocui.ModNone, nba.keyfn(nba.toggleGameBoxScore)); err != nil {
	//	return err
	//}

	if err := g.SetKeybinding(scoreboardLabel, gocui.KeyCtrlT, gocui.ModNone, nba.keyfn(nba.toggleTeamStats)); err != nil {
		return err
	}

	return nil
}

func (nba *NBAConsole) moveDown() error {
	_, y := nba.g.CurrentView().Cursor()
	if len(nba.gamesList.gameIDs)+1 /** TODO: ugh (len(games)- 1) + 2 = + 1*/ > y {
		nba.g.CurrentView().SetCursor(0, y+1) // TODO: limit scrolling down
		go nba.toggleGameBoxScore()
	}

	return nil
}

func (nba *NBAConsole) moveUp() error {
	_, y := nba.g.CurrentView().Cursor()
	if y > 2 { // 2 is the magic num
		nba.g.CurrentView().SetCursor(0, y-1)
		go nba.toggleGameBoxScore()
	}

	return nil
}

func (nba *NBAConsole) keyfn(fn func() error) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return fn()
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
