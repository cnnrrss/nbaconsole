package app

import (
	"fmt"
	"time"
)

var (
	footerFmt            string = "v%s  %s  %s  %s  %s[%s]"
	quitText             string = "Quit[Q]"
	lastUpdatedText      string = "Updated"
	toggleBoxScoreText   string = "Stats[Ctrl+t]"
	toggleScoreBoardText string = "Scores[Ctrl+r]"
)

func (nba *NBAConsole) updateFooter(s string /** TODO */) {
	nba.update(func() {
		nba.footerView.Clear()
		nba.lastUpdated = time.Now()
		fmt.Fprintln(nba.footerView,
			fmt.Sprintf(footerFmt,
				nba.version(),
				quitText,
				toggleBoxScoreText,
				toggleScoreBoardText,
				lastUpdatedText,
				toHumanTime(nba.lastUpdated),
			),
		)
	})
}
