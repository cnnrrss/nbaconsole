package app

import (
	"fmt"
	"time"

	"github.com/cnnrrss/nbaconsole/common/pad"
)

var (
	versionText     string = "Version:"
	quitText        string = "Quit: [Q]"
	lastUpdatedText string = "Last Updated:"
	toggleBoxScoreText string = "BoxScoreView: [Ctrl+t]"
)

func (nba *NBAConsole) updateFooter(s string) {
	nba.update(func() {
		nba.footerView.Clear()
		nba.lastUpdated = time.Now()
		fmt.Fprintln(nba.footerView,
			fmt.Sprintf("%s [%s] %s %s %s %s",
				versionText,
				nba.version(),
				pad.Left(quitText, 8, " "),
				lastUpdatedText,
				toHumanTime(nba.lastUpdated),
				toggleBoxScoreText,
			),
		)
	})
}
