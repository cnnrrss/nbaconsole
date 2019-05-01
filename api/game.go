package api

import (
	"fmt"
	"strings"
	"time"
)

// IsActive returns whether the game is still ongoing
func (g *Game) IsActive() bool {
	if g.StatusNum == 3 {
		return false
	}
	return true
}

// HasStarted returns whether the game has started
func (g *Game) HasStarted() bool {
	if g.StartTimeUTC.Before(time.Now().UTC()) {
		return false
	}
	return true
}

// Status returns whether the game has started
func (g *Game) Status() string {
	currTime := time.Now().UTC()
	if g.IsGameActivated {
		return fmt.Sprintf("%s Q%d", g.Clock, g.Period.Current)
	} else if g.StatusNum == 3 && g.StartTimeUTC.Before(currTime) {
		return "Final"
	}
	return fmt.Sprintf("%v", g.StartTimeEastern)
}

// Summary returns nugget of info about the game
func (g *Game) Summary() string {
	return g.Nugget.Text
}

// IsOvertime returns the number of overtime periods in the game
func (g *Game) IsOvertime() bool {
	return g.Period.Current > 4
}

// IsPlayoffs returns a the winning team series note
func (g *Game) IsPlayoffs() string {
	return g.Playoffs.SeriesSummaryText
}

// Score returns the score of the game in two strings: home, visitor
func (g *Game) Score() (home, visitor string) {
	return g.HTeam.Score, g.VTeam.Score
}

func playOffRound(r int, div string) string {
	switch r {
	case 1:
		return fmt.Sprintf("%sern Conference Quarter Finals", strings.Title(div))
	case 2:
		return fmt.Sprintf("%sern Conference Semi Finals", strings.Title(div))
	case 3:
		return fmt.Sprintf("%sern Conference Finals", strings.Title(div))
	case 4:
		return fmt.Sprintf("Finals")
	default:
		return ""
	}
}
