package app

import (
	"github.com/jroimartin/gocui"
)

func highlightView(v *gocui.View) {
	// v.SetCursor(0, y+2)
	v.Highlight = true
	v.SelFgColor = gocui.ColorBlue
	v.SelBgColor = gocui.ColorDefault
}
