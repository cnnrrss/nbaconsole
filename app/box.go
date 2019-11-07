package app

import (
	"fmt"

	"github.com/cnnrrss/nbaconsole/api"
	"github.com/jroimartin/gocui"
)

// TODO: v2 API, handle multipl pages

// Page ...
type Page struct {
	Offset, Limit int
}

// Box extends gocui.View with game specific info
type Box struct {
	*gocui.View
	games      []string
	pages      []Page
	currentIdx int
	sorted     bool
}

type GameScore struct {
	api.GameBoxScore
	// ID of the GameScore used to make
	// additional requests to the NBA API
	ID string
}

// NewBox initializes a Box object with an existing gocui.View
func NewBox(v *gocui.View, sorted bool) *Box {
	b := &Box{}
	b.View = v
	b.SelBgColor = gocui.ColorBlack
	b.SelFgColor = gocui.ColorWhite
	b.Autoscroll = true
	b.sorted = sorted

	return b
}

// Wipe wipes a box from the terminal
func (b *Box) Wipe() {
	b.games = make([]string, 0)
	b.pages = []Page{}
	b.Clear()
	b.SetCursor(0, 0)
}

// IsEmpty indicates whether a Box has any items
func (b *Box) IsEmpty() bool {
	return len(b.games) == 0
}

// SetHeader will set the title of the View and display paging information of the
// list if there are more than one pages
func (b *Box) SetHeader(header string) {
	b.Title = header

	if b.pagesNum() > 1 {
		b.Title = fmt.Sprintf(" %d/%d - %v ", b.currPageNum(), b.pagesNum(), header)
	} else {
		b.Title = fmt.Sprintf(" %v ", header)
	}
}

func (b *Box) currPageNum() int {
	if b.IsEmpty() {
		return 0
	}
	return b.currentIdx + 1
}

func (b *Box) pagesNum() int {
	return len(b.pages)
}

// Focus hightlights the View of the current List
// func (b *Box) Focus(g *gocui.Gui) error {
// 	b.Highlight = true
// 	_, err := g.SetCurrentView(b.Name())

// 	return err
// }

// func (b *Box) currPage() Page {
// 	return b.pages[b.currentIdx]
// }

// func (b *Box) displayGame(bi int) error {
// 	b.Clear()
// 	b.currentIdx = bi
// 	box := b.pages[b.currentIdx]
// 	for i := box.Offset; i < box.Offset+box.Limit; i++ {
// 		if _, err := fmt.Fprintln(b.View, b.displayBoxGame(i)); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (b *Box) displayBoxGame(i int) string {
// 	item := fmt.Sprint(b.games[i])
// 	x, _ := b.Size()
// 	sp := pad.AddString(x - 1 - len(item) - 3, " ")
// 	if b.sorted {
// 		return fmt.Sprintf(">>>>%2d. %v%v", i+1, item, sp)
// 	}
// 	return fmt.Sprintf(">>> %v%v", item, sp)
// }
