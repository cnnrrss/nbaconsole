package model

import (
	"github.com/connorvanderhook/nbaconsole/api"
	"github.com/jroimartin/gocui"
)

// NBAConsole provides the context for running the app
type NBAConsole struct {
	g              *gocui.Gui
	scoreboardView *gocui.View
	client         *api.Client
	debug          bool
}

// Config options
type Config struct {
	APIKey string
}
