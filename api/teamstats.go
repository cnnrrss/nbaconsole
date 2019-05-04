package api

import (
	"fmt"
	"strings"

	"strconv"
)

var (
	lineTeamStatsHeaderFmt = "%-8s%-6s%7s\n"
	lineShotAttemptsFmt    = "%-9s%-7s%5s\n"
	lineShotPercentageFmt  = "%-10s%-7s%3s\n"
	lineReboundStatsFmt    = "%-7d%s%7d\n"
	lineAssistStatsFmt     = "%-7s%s%7s\n"
	lineStealStatsFmt      = "%-8s%s%7s\n"
	lineTurnoverStatsFmt   = "%-6s%-6s%6s\n"
)

func (bs *GameBoxScore) TeamStats() string {
	var str strings.Builder
	if bs != nil {
		home, away := bs.SportsContent.Game.Home, bs.SportsContent.Game.Visitor
		str.WriteString(
			fmt.Sprintf(lineTeamStatsHeaderFmt,
				home.Abbreviation,
				"STATS",
				away.Abbreviation,
			),
		)
		str.WriteString(
			fmt.Sprintf(lineShotAttemptsFmt,
				toFraction(home.Stats.FieldGoalsMade, home.Stats.FieldGoalsAttempted),
				"FGs",
				toFraction(away.Stats.FieldGoalsMade, away.Stats.FieldGoalsAttempted),
			),
		)
		str.WriteString(
			fmt.Sprintf(lineShotPercentageFmt,
				home.Stats.FieldGoalsPercentage,
				"%",
				away.Stats.FieldGoalsPercentage,
			),
		)
		str.WriteString(
			fmt.Sprintf(lineShotAttemptsFmt,
				toFraction(home.Stats.ThreePointersMade, home.Stats.ThreePointersAttempted),
				"3pts",
				toFraction(away.Stats.ThreePointersMade, away.Stats.ThreePointersAttempted),
			),
		)
		str.WriteString(
			fmt.Sprintf(lineShotPercentageFmt,
				home.Stats.ThreePointersPercentage,
				"%",
				away.Stats.ThreePointersPercentage,
			),
		)
		str.WriteString(
			fmt.Sprintf(lineShotAttemptsFmt,
				toFraction(home.Stats.FreeThrowsMade, home.Stats.FreeThrowsAttempted),
				"FTs",
				toFraction(away.Stats.FreeThrowsMade, away.Stats.FreeThrowsAttempted),
			),
		)
		str.WriteString(
			fmt.Sprintf(lineShotPercentageFmt,
				home.Stats.FreeThrowsPercentage,
				"%",
				away.Stats.FreeThrowsPercentage,
			),
		)
		homeDefReb, _ := strconv.Atoi(home.Stats.ReboundsDefensive)
		awayDefReb, _ := strconv.Atoi(away.Stats.ReboundsDefensive)
		homeOffReb, _ := strconv.Atoi(home.Stats.ReboundsOffensive)
		awayOffReb, _ := strconv.Atoi(away.Stats.ReboundsOffensive)
		str.WriteString(
			fmt.Sprintf(lineReboundStatsFmt,
				homeDefReb+homeOffReb,
				"Tot Reb",
				awayDefReb+awayOffReb,
			),
		)
		str.WriteString(
			fmt.Sprintf(lineReboundStatsFmt,
				homeOffReb,
				"Off Reb",
				awayOffReb,
			),
		)
		str.WriteString(
			fmt.Sprintf(lineAssistStatsFmt,
				home.Stats.Assists,
				"Assists",
				away.Stats.Assists,
			),
		)
		str.WriteString(
			fmt.Sprintf(lineStealStatsFmt,
				home.Stats.Steals,
				"Steals",
				away.Stats.Steals,
			),
		)
		str.WriteString(
			fmt.Sprintf(lineTurnoverStatsFmt,
				home.Stats.Turnovers,
				"Turnovers",
				away.Stats.Turnovers,
			),
		)

	} else {
		str.WriteString("error getting team stats")
	}
	return str.String()
}

func toFraction(a, b string) string {
	return fmt.Sprintf("%s/%s", a, b)
}
