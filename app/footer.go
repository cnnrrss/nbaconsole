package app

import "fmt"

var (
	quitText = "Quit"
)

func (nba *NBAConsole) updateFooter(s string) {
	nba.update(func() {
		nba.footerView.Clear()
		base := fmt.Sprintf("%s: [Q]", quitText)
		v := fmt.Sprintf(" Version: [%s]", nba.version())
		str := PadLeft(fmt.Sprintf("%s", base), 8)
		fmt.Fprintln(nba.footerView, fmt.Sprintf("%s %s", v, str))
	})
}
