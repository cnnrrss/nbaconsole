package nbaconsole

import (
	"github.com/jroimartin/gocui"
	log "github.com/sirupsen/logrus"
)

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, quit); err != nil {
		return err
	}

	return nil
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

func (nba *NBAConsole) keyfn(fn func() error) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return fn()
	}
}
