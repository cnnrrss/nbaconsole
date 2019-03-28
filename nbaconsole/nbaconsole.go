package nbaconsole

import (
	"fmt"
	"log"
	"os"

	"github.com/jroimartin/gocui"
)

// NBAConsole provides the context for running the app
type NBAConsole struct {
	g     *gocui.Gui
	debug bool
}

// Config options
type Config struct {
	APIKey string
}

// NewNBAConsole loads a new context for running the app
func NewNBAConsole(config *Config) *NBAConsole {
	var debug bool
	if os.Getenv("DEBUG") != "" {
		debug = true
	}

	return &NBAConsole{
		debug: debug,
	}
}

// Start the NBA Console
func (nba *NBAConsole) Start() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatalf("new gocui: %v", err)
	}
	nba.g = g
	defer g.Close()

	g.InputEsc = true
	g.FgColor = gocui.ColorWhite
	g.Mouse = true
	g.Highlight = true

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	// g.SetManagerFunc(nba.layout)
	// if err := ct.keybindings(g); err != nil {
	// 	log.Fatalf("keybindings: %v", err)
	// }

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalf("main loop: %v", err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", maxX/2-13, maxY/2-5, maxX/2+13, maxY/2+5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, nbaConsoleLogo)
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
