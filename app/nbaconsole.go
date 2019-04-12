package nbaconsole

import (
	"fmt"
	"log"
	"os"

	"github.com/connorvanderhook/nbaconsole/api"
	"github.com/jroimartin/gocui"
)

// TODO: organize types in proj

// NBAConsole provides the context for running the app
type NBAConsole struct {
	g              *gocui.Gui
	scoreboardView *gocui.View
	client         *api.Client
	debug          bool
}

// NewNBAConsole loads a new context for running the app
func NewNBAConsole(config *Config) *NBAConsole {
	var debug bool
	if os.Getenv("DEBUG") != "" {
		debug = true
	}

	client := api.NewClient()

	return &NBAConsole{
		debug:  debug,
		client: client,
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

	/* --------------------------------------------------- */
	// TODO refactor to nba.Settings()
	g.InputEsc = true
	g.Mouse = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorBlue
	g.BgColor = gocui.ColorBlack
	g.FgColor = gocui.ColorWhite
	/* --------------------------------------------------- */

	// The terminal’s width and height are needed for layout calculations.
	tWidth, tHeight := g.Size()
	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln("Could not set key binding:", err)
	}
	/* --------------------------------------------------- */
	// TODO: refactor to use nba.scoreBoardLayout()
	scoreboardView, err := g.SetView(scoreBoardLayoutName, 0, 0, tWidth-4, tHeight-4)
	if err != nil && err != gocui.ErrUnknownView {
		log.Println("Failed to create scoreboardView:", err)
		return
	}

	scoreboardView.FgColor = gocui.ColorMagenta
	// Let the view scroll if the output exceeds the visible area.
	scoreboardView.Autoscroll = true
	scoreboardView.Wrap = true
	nba.scoreboardView = scoreboardView

	nba.getScoreboard()
	/* --------------------------------------------------- */

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalf("main loop: %v", err)
	}
}

func layout(g *gocui.Gui) error {
	// get current terminal size
	maxX, maxY := g.Size()
	if v, err := g.SetView("Welcome", maxX/2-13, maxY/2-5, maxX/2+13, maxY/2+5); err != nil {
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
