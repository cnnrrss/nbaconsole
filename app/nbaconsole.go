package nbaconsole

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/connorvanderhook/nbaconsole/api"
	"github.com/jroimartin/gocui"
)

// NBAConsole provides the context for running the app
type NBAConsole struct {
	g              *gocui.Gui
	scoreboardView *gocui.View
	client         *api.Client
	// TODO: don't hardcode date
	date          string
	forceRefresh  chan bool
	refreshTicker *time.Ticker
	rateLimiter   <-chan time.Time
	debug         bool
}

// NewNBAConsole loads a new context for running the app
func NewNBAConsole() *NBAConsole {
	var debug bool
	if os.Getenv("DEBUG") != "" {
		debug = true
	}

	client := api.NewClient()
	curDate := currentDate()

	return &NBAConsole{
		client:        client,
		date:          curDate,
		forceRefresh:  make(chan bool),
		refreshTicker: time.NewTicker(1 * time.Minute),
		rateLimiter:   time.Tick(10 * time.Second),
		debug:         debug,
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
	if v, err := g.SetView("Welcome", 0, 0, 40, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		fmt.Fprintf(v, "%5s\n", "Welcome to NBA Console")
	}
	return nil
}

func (nba *NBAConsole) refresh() error {
	go func() {
		<-nba.rateLimiter
		nba.forceRefresh <- true
	}()
	return nil
}

func (nba *NBAConsole) pollScoreboardData() {
	go func() {
		for {
			select {
			case <-nba.forceRefresh:
				nba.refreshAll()
			case <-nba.refreshTicker.C:
				nba.refreshAll()
			}
		}
	}()
}

func (nba *NBAConsole) refreshAll() error {
	// TODO: do you need mutex or cache here?
	go func() {
		fmt.Fprintf(nba.scoreboardView, "refreshing...")
		nba.getScoreboard()
	}()

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
