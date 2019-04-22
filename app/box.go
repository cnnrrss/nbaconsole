package app

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// Box extends gocui.View with game specific info
type Box struct {
	*gocui.View
	Header   string
	games    []interface{}
	pages    []Page
	curIndex int
	ordered  bool
}

// NewBox initializes a Box object with an existing gocui.View
func NewBox(v *gocui.View, ordered bool) *Box {
	b := &Box{}
	b.View = v
	b.SelBgColor = gocui.ColorBlack
	b.SelFgColor = gocui.ColorWhite
	b.Autoscroll = true
	b.ordered = ordered

	return b
}

// Wipe wipes a box from the terminal
func (b *Box) Wipe() {
	b.games = make([]interface{}, 0)
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
	b.Header = header

	if b.pagesNum() > 1 {
		b.Header = fmt.Sprintf(" %d/%d - %v ", b.currPageNum(), b.pagesNum(), header)
	} else {
		b.Header = fmt.Sprintf(" %v ", header)
	}
}

// Focus hightlights the View of the current List
func (b *Box) Focus(g *gocui.Gui) error {
	b.Highlight = true
	_, err := g.SetCurrentView(b.Name())

	return err
}

// currPageNum returns the current page number to display
func (b *Box) currPageNum() int {
	if b.IsEmpty() {
		return 0
	}
	return b.curIndex + 1
}

// pageNum returns the number of the pages
func (b *Box) pagesNum() int {
	return len(b.pages)
}

// currPage returns the current page being displayd
func (b *Box) currPage() Page {
	return b.pages[b.curIndex]
}
