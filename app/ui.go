package app

import (
	"github.com/jroimartin/gocui"
)

func highlightView(v *gocui.View) {
	v.Highlight = true
	v.SelFgColor = gocui.ColorBlue
	v.SelBgColor = gocui.ColorDefault
}
