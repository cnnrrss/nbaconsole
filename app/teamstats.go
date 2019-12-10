package app

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func (nba *NBAConsole) toggleTeamStats() error {
	currentGameID := nba.getSelectedGameID()
	if nba.selectedGameScore.GameID() != currentGameID {
		nba.selectedGameID = currentGameID
	}

	go nba.setTeamStatsView(nba.g)

	return nil
}

func (nba *NBAConsole) setTeamStatsView(g *gocui.Gui) error {
	if v, err := g.SetView(
		teamStatsLabel,
		0, /** globalX0 */
		0, /** scoreboardY0 */
		nba.curW-1,
		nba.curH-footerHeight-footerHeight,
	); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		nba.teamStats = v
		nba.teamStats.Frame = false

		go func() {
			nba.setTeamStats()
		}()
	}
	return nil
}

func (nba *NBAConsole) setTeamStats() error { // TODO: change to draw for consistent API
	nba.g.SetCurrentView(teamStatsLabel)
	nba.update(func() {
		nba.teamStats.Clear()
		fmt.Fprintln(nba.teamStats, fmt.Sprintf("%s\n\n", nba.selectedGameScore.TeamStats()))
		fmt.Fprintln(nba.teamStats, nba.selectedGameScore.BoxScoreLeaders())
	})
	return nil
}
