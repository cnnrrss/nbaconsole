package nbaconsole

import (
	"github.com/jroimartin/gocui"
)

// relSize returns window width and height
func relSize(g *gocui.Gui) (int, int) {
	tw, th := g.Size()
	return tw, th
}
