package app

import (
	"fmt"
	"time"

	"github.com/cnnrrss/nbaconsole/common/pad"
)

var (
	quitText = "Quit"
)

func (nba *NBAConsole) updateFooter(s string) {
	nba.update(func() {
		nba.footerView.Clear()
		base := fmt.Sprintf("%s: [Q]", quitText)
		v := fmt.Sprintf(" Version: [%s]", nba.version())
		str := pad.Left(base, 8, " ")
		nba.lastUpdated = time.Now()
		upd := fmt.Sprintf("Last Updated: [%s]", toHumanTime(nba.lastUpdated))
		fmt.Fprintln(nba.footerView, fmt.Sprintf("%s %s %s", v, str, upd))
	})
}
