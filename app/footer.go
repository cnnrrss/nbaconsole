package nbaconsole

import "fmt"

func (nba *NBAConsole) updateFooter(s string) {
	var quitText string
	quitText = "Quit"

	nba.update(func() {
		nba.footerView.Clear()
		base := fmt.Sprintf("%s: [Q]\n", quitText)
		v := fmt.Sprintf("v%s", nba.version())
		str := PadRight(fmt.Sprintf("%s", base), len(base))
		str = str[:len(str)-len(v)+2] + v
		fmt.Fprintln(nba.footerView, str)
	})
}
