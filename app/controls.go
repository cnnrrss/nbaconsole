package app

import (
	"github.com/jroimartin/gocui"
)

func (nba *NBAConsole) keybindings(g *gocui.Gui) error {
	var err error
	if err = g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, MoveUp); err != nil {
		return err
	}

	if err = g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, MoveDown); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding(scoreboardLabel, gocui.KeyEnter, gocui.ModNone, nba.keyfn(nba.ToggleGameBoxScore)); err != nil {
		return err
	}

	if err := g.SetKeybinding(scoreboardLabel, gocui.KeyCtrlT, gocui.ModNone, nba.keyfn(nba.ToggleTeamStats)); err != nil {
		return err
	}

	// TODO:
	if err := g.SetKeybinding("teamstats", gocui.KeyEnter, gocui.ModNone, nba.keyfn(nba.getScoreboard)); err != nil {
		return err
	}

	return nil
}

func (nba *NBAConsole) keyfn(fn func() error) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return fn()
	}
}

// MoveDown modifies the gocui cursor +1 on y axis
func MoveDown(g *gocui.Gui, v *gocui.View) error {
	_, y := g.CurrentView().Cursor()
	v.SetCursor(0, y+1)
	return nil
}

// MoveUp modifies the gocui cursor +1 on y axis
func MoveUp(g *gocui.Gui, v *gocui.View) error {
	_, y := g.CurrentView().Cursor()
	v.SetCursor(0, y-1)
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
