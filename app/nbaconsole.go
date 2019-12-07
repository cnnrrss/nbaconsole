package app

import (
	"log"
	"math/rand"
	"time"

	"github.com/jroimartin/gocui"
)

// NBAConsole provides the context for running the app
type NBAConsole struct {
	// gocui User Interface
	// TODO: unwrap maybe?
	g *gocui.Gui

	// Views TODO: make array or map?
	footerView *gocui.View
	scoreboard *gocui.View
	boxScore   *gocui.View
	teamStats  *gocui.View
	helpView   *gocui.View

	// refresh ticker
	// TODO: implement done and rate limiting
	refreshTicker *time.Ticker
	rateLimiter   <-chan time.Time
	lastUpdated   time.Time
	forceRefresh  chan bool
	done          chan bool

	// stateful nba game data
	selectedGame      string
	selectedGameScore *GameScore // TODO: implement caching
	gamesList         *Box

	// additional console state
	message string
	date    string
	debug   bool
	curW    int
	curH    int
}

// NewNBAConsole loads a new context for running the app
func NewNBAConsole(date, tz string, debug bool) *NBAConsole {
	if date == "" {
		date = currentDateYYYYMMDD(tz)
	}

	return &NBAConsole{
		date:          date,
		debug:         debug,
		message:       nbaMessages[rand.Intn(len(nbaMessages)-1)], // generate random hello
		forceRefresh:  make(chan bool),
		refreshTicker: time.NewTicker(30 * time.Second),
		rateLimiter:   time.Tick(60 * time.Second),
		lastUpdated:   time.Now(),
	}
}

// Start the NBA Console
func (nba *NBAConsole) Start() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatalf("Failed to initialize new gocui: %v", err)
	}
	defer g.Close()

	nba.g = g
	gw, gh := g.Size()
	nba.curW, nba.curH = min(gw, MINWIDTH), min(gh, MINHEIGHT)
	g.SetManagerFunc(nba.layout)
	nba.keybindings(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalf("main loop exiting: %v", err)
	}
	log.Println("Exiting")
}
