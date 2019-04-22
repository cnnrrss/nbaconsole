package app

import (
	"bytes"
	"fmt"

	"github.com/jroimartin/gocui"
)

// TODO: v2 API, handle multipl pages

// Page ...
type Page struct {
	Offset, Limit int
}

// Box extends gocui.View with game specific info
type Box struct {
	*gocui.View  // inline a gocui View
	Title        string
	Items        []interface{}
	Pages        []Page
	CurrentIndex int
	Ordered      bool
}

// NewBox initializes a Box object with an existing gocui.View
func NewBox(v *gocui.View, ordered bool) *Box {
	b := &Box{}
	b.View = v
	b.SelBgColor = gocui.ColorBlack
	b.SelFgColor = gocui.ColorWhite
	b.Autoscroll = true
	b.Ordered = ordered

	return b
}

// Wipe wipes a box from the terminal
func (b *Box) Wipe() {
	b.Items = make([]interface{}, 0)
	b.Pages = []Page{}
	b.Clear()
	b.SetCursor(0, 0)
}

// IsEmpty indicates whether a Box has any items
func (b *Box) IsEmpty() bool {
	return len(b.Items) == 0
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
	return b.CurrentIndex + 1
}

// pageNum returns the number of the pages
func (b *Box) pagesNum() int {
	return len(b.Pages)
}

// currPage returns the current page being displayd
func (b *Box) currPage() Page {
	return b.Pages[b.CurrentIndex]
}

func (b *Box) displayBox(bi int) error {
	b.Clear()           // clear existing view
	b.CurrentIndex = bi // set pageIdx
	box := b.Pages[b.CurrentIndex]
	for i := box.Offset; i < box.Offset+box.Limit; i++ {
		if _, err := fmt.Fprintln(b.View, b.displayBoxItem(i)); err != nil {
			return err
		}
	}
	return nil
}

// displayBoxItem ...
func (b *Box) displayBoxItem(i int) string {
	item := fmt.Sprint(b.Items[i])
	x, _ := b.Size()
	sp := addSpaces(x - 1 - len(item) - 3)
	if b.Ordered {
		return fmt.Sprintf(">>>>%2d. %v%v", i+1, item, sp)
	}
	return fmt.Sprintf(">>> %v%v", item, sp)
}

func addSpaces(n int) string {
	var s bytes.Buffer
	for i := 0; i < n; i++ {
		s.WriteString(" ")
	}
	return s.String()
}
