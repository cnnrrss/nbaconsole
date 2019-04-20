package nbaconsole

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jroimartin/gocui"
)

var nbaDate string

// NBAConsole provides the context for running the app
type NBAConsole struct {
	g             *gocui.Gui
	headers       *gocui.View
	footerView    *gocui.View
	scoreboard    *gocui.View
	helpView      *gocui.View
	gamesList     *Box
	date          string
	forceRefresh  chan bool
	done          chan bool
	refreshTicker *time.Ticker
	rateLimiter   <-chan time.Time
	debug         bool
	curW          int
	curH          int
}

func init() {
	flag.StringVar(&nbaDate, "d", "", "optionally retrieve NBA scoreboard for date in YYYYMMDD format")
	flag.Parse()
}

// NewNBAConsole loads a new context for running the app
func NewNBAConsole() *NBAConsole {
	var debug bool
	if os.Getenv("DEBUG") != "" {
		debug = true
	}

	if nbaDate == "" {
		nbaDate = currentDate()
	}

	return &NBAConsole{
		date:          nbaDate,
		forceRefresh:  make(chan bool),
		refreshTicker: time.NewTicker(30 * time.Second),
		rateLimiter:   time.Tick(10 * time.Second),
		debug:         debug,
	}
}

// Start the NBA Console
func (nba *NBAConsole) Start() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatalf("Failed to initialize new gocui: %v", err)
	}
	nba.g = g
	nba.curW, nba.curH = g.Size()
	defer g.Close()

	g.InputEsc = true
	g.Mouse = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorBlue
	g.BgColor = gocui.ColorBlack
	g.FgColor = gocui.ColorWhite

	g.SetManagerFunc(nba.layout)

	if err = g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, MoveUp); err != nil {
		log.Fatal("Could not set keybinding:", err)
	}

	if err = g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, MoveDown); err != nil {
		log.Fatal("Could not set keybinding:", err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Fatal("Could not set keybinding:", err)
	}

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, nil); err != nil {
		log.Fatal("Could not set keybinding:", err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalf("main loop exiting: %v", err)
	}
	log.Println("Exiting")
}

func (nba *NBAConsole) refresh() error {
	go func() {
		<-nba.rateLimiter
		nba.forceRefresh <- true
	}()
	return nil
}

func (nba *NBAConsole) pollScoreboardData(params map[string]string) {
	go func() {
		for {
			select {
			case <-nba.forceRefresh:
				nba.refreshAll(params)
			case <-nba.refreshTicker.C:
				nba.refreshAll(params)
			}
		}
	}()
}

func (nba *NBAConsole) refreshAll(params map[string]string) error {
	go func() {
		fmt.Fprintf(nba.scoreboard, "refreshing...")
		nba.scoreboard.Clear()
		nba.getScoreboard(params)
	}()
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
