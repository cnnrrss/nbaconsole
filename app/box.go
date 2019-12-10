package app

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// Page ...
type Page struct {
	Offset, Limit int
}

// Box extends gocui.View with game specific info
type Box struct {
	*gocui.View
	gameIDs    []string
	pages      []Page
	currentIdx int
	sorted     bool
}

// newBox initializes a Box object with an existing gocui.View
func newBox(v *gocui.View, sorted bool) *Box {
	b := &Box{}
	b.View = v
	b.SelBgColor = gocui.ColorBlack
	b.SelFgColor = gocui.ColorWhite
	b.Autoscroll = true
	b.sorted = sorted
	return b
}

// wipe wipes a box from the terminal
func (b *Box) wipe() {
	b.gameIDs = make([]string, 0)
	b.pages = []Page{}
	b.Clear()
	b.SetCursor(0, 2)
}

// isEmpty indicates whether a Box has any items
func (b *Box) isEmpty() bool {
	return len(b.gameIDs) == 0
}

// setHeader will set the title of the View and display paging information of the
// list if there are more than one pages
func (b *Box) setHeader(header string) {
	b.Title = header

	if b.pagesNum() > 1 {
		b.Title = fmt.Sprintf(" %d/%d - %v ", b.currPageNum(), b.pagesNum(), header)
	} else {
		b.Title = fmt.Sprintf(" %v ", header)
	}
}

func (b *Box) currPageNum() int {
	if b.isEmpty() {
		return 0
	}
	return b.currentIdx + 1
}

func (b *Box) pagesNum() int {
	return len(b.pages)
}
