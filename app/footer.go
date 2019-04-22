package app

import (
	"fmt"

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
		str := pad.Left(fmt.Sprintf("%s", base), 8, " ")
		fmt.Fprintln(nba.footerView, fmt.Sprintf("%s %s", v, str))
	})
}
