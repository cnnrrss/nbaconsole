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

func (nba *NBAConsole) keyfn(fn func() error) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return fn()
	}
}
