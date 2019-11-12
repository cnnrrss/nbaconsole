package app

import (
	"fmt"
	"time"
)

var (
	footerFmt            string = "%s %s %s %s"
	quitText             string = "Quit[Q]"
	toggleBoxScoreText   string = "Box[enter]"
	toggleTeamStatsText  string = "Stats[Ctl+t]"
	toggleScoreBoardText string = "Home[Ctl+r]"
)

func (nba *NBAConsole) updateFooter(s string /** TODO */) {
	nba.update(func() {
		nba.footerView.Clear()
		nba.lastUpdated = time.Now()
		fmt.Fprintln(nba.footerView,
			fmt.Sprintf(footerFmt,
				quitText,
				toggleBoxScoreText,
				toggleTeamStatsText,
				toggleScoreBoardText,
			),
		)
	})
}
