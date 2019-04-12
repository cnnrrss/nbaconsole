package nbaconsole

import (
	"github.com/jroimartin/gocui"
	log "github.com/sirupsen/logrus"
)

// update update view
func (nba *NBAConsole) update(f func()) {
	if nba.g == nil {
		log.Fatal("gocui is not initialized")
	}

	nba.g.Update(func(g *gocui.Gui) error {
		f()
		return nil
	})
}
