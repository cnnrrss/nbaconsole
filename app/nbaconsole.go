package app

import (
	"log"
	"math/rand"
	"time"

	"github.com/jroimartin/gocui"
)

// NBAConsole provides the context for running the app
type NBAConsole struct {
	g             *gocui.Gui
	headers       *gocui.View
	footerView    *gocui.View
	scoreboard    *gocui.View
	boxScore      *gocui.View
	helpView      *gocui.View
	Views         map[string]*gocui.View
	selectedGame  *GameScore
	gamesList     *Box
	refreshTicker *time.Ticker
	rateLimiter   <-chan time.Time
	forceRefresh  chan bool
	done          chan bool
	message       string
	date          string
	debug         bool
	curW          int
	curH          int
}

// NewNBAConsole loads a new context for running the app
func NewNBAConsole(date string, debug bool) *NBAConsole {
	if date == "" {
		date = currentDate()
	}

	return &NBAConsole{
		date:          date,
		debug:         debug,
		message:       nbaMessages[rand.Intn(len(nbaMessages)-1)], // generate random hello
		forceRefresh:  make(chan bool),
		refreshTicker: time.NewTicker(60 * time.Second),
		rateLimiter:   time.Tick(10 * time.Second),
	}
}

// Start the NBA Console
func (nba *NBAConsole) Start() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatalf("Failed to initialize new gocui: %v", err)
	}
	defer g.Close()

	g.InputEsc = true
	g.Mouse = true
	g.Highlight = true
	g.Cursor = true
	g.BgColor = gocui.ColorBlack
	g.FgColor = gocui.ColorWhite

	nba.g = g
	nba.curW, nba.curH = g.Size()
	g.SetManagerFunc(nba.layout)
	nba.keybindings(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalf("main loop exiting: %v", err)
	}
	log.Println("Exiting")
}
