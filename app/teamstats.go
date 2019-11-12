package app

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// ToggleGameBoxScore toggles between the global scoreboard and the game box score
func (nba *NBAConsole) ToggleTeamStats() error {
	selectedGame := nba.SelectedGame()
	if nba.selectedGame != selectedGame {
		nba.selectedGame = selectedGame
	}

	go nba.setTeamStatsView(nba.g, nba.selectedGame)

	return nil
}

func (nba *NBAConsole) setTeamStatsView(g *gocui.Gui, gameID string) error {
	if v, err := g.SetView(teamStatsLabel, 0 /** globalX0 */, 0 /** scoreboardY0 */, nba.curW-1, nba.curH-footerHeight-footerHeight); err != nil {
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

func (nba *NBAConsole) setTeamStats() error {
	nba.g.SetCurrentView(teamStatsLabel)
	nba.update(func() {
		nba.teamStats.Clear()
		fmt.Fprintln(nba.teamStats, nba.selectedGameScore.TeamStats())
	})
	return nil
}
