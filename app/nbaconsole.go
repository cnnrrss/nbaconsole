package nbaconsole

import (
	"fmt"
	"log"
	"os"

	"github.com/connorvanderhook/nbaconsole/api"
	"github.com/jroimartin/gocui"
	"github.com/pkg/errors"
)

// NBAConsole provides the context for running the app
type NBAConsole struct {
	g              *gocui.Gui
	scoreboardView *gocui.View
	client         *api.Client
	date           string
	debug          bool
}

// NewNBAConsole loads a new context for running the app
func NewNBAConsole(config *Config) *NBAConsole {
	var debug bool
	if os.Getenv("DEBUG") != "" {
		debug = true
	}

	client := api.NewClient()
	curDate := currentDate()

	return &NBAConsole{
		debug:  debug,
		client: client,
		date:   curDate,
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
	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln("Could not set key binding:", err)
	}

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, nil); err != nil {
		log.Panicln("Could not set key binding:", err)
	}

	if err := nba.scoreBoardLayout(g); err != nil {
		log.Println(err)
		return
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalf("main loop: %v", err)
	}
}

func layout(g *gocui.Gui) error {
	// get current terminal size
	tWidth, tHeight := g.Size()
	if v, err := g.SetView("Welcome", tWidth/2-13, tHeight/2-5, tWidth/2+13, tHeight/2+5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, nbaConsoleLogo)
	}

	// Scoreboard
	_, err := g.SetView(scoreBoardLayoutName, 22, 0, tWidth, tHeight-1)
	if err != nil {
		return errors.Wrap(err, "Cannot update tasks view")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
