package api

import (
	"fmt"
	"strings"

	"strconv"

	"github.com/cnnrrss/nbaconsole/common/pad"
)

func (bs *GameBoxScore) TeamStats() string {
	var str strings.Builder
	if bs != nil {
		home, away := bs.SportsContent.Game.Home, bs.SportsContent.Game.Visitor
		str.WriteString(fmt.Sprintf("%s%s%s%s%s\n", home.Abbreviation, pad.AddSpaces(5), "STATS", pad.AddSpaces(5), away.Abbreviation))
		str.WriteString(
			fmt.Sprintf("%s/%s%s%s%s%s/%s\n",
				home.Stats.FieldGoalsMade,
				home.Stats.FieldGoalsAttempted,
				pad.AddSpaces(4),
				"FGs",
				pad.AddSpaces(4),
				away.Stats.FieldGoalsMade,
				away.Stats.FieldGoalsAttempted,
			),
		)
		str.WriteString(
			fmt.Sprintf("%s%s%s%s%s\n",
				home.Stats.FieldGoalsPercentage,
				pad.AddSpaces(6),
				"%",
				pad.AddSpaces(6),
				away.Stats.FieldGoalsPercentage,
			),
		)
		str.WriteString(
			fmt.Sprintf("%s/%s%s%s%s%s/%s\n",
				home.Stats.ThreePointersMade,
				home.Stats.ThreePointersAttempted,
				pad.AddSpaces(4),
				"3pts",
				pad.AddSpaces(4),
				away.Stats.ThreePointersMade,
				away.Stats.ThreePointersAttempted,
			),
		)
		str.WriteString(
			fmt.Sprintf("%s%s%s%s%s\n",
				home.Stats.ThreePointersPercentage,
				pad.AddSpaces(6),
				"%",
				pad.AddSpaces(6),
				away.Stats.ThreePointersPercentage,
			),
		)
		str.WriteString(
			fmt.Sprintf("%s/%s%s%s%s%s/%s\n",
				home.Stats.FreeThrowsMade,
				home.Stats.FreeThrowsAttempted,
				pad.AddSpaces(4),
				"FTs",
				pad.AddSpaces(4),
				away.Stats.FreeThrowsMade,
				away.Stats.FreeThrowsAttempted,
			),
		)
		str.WriteString(
			fmt.Sprintf("%s%s%s%s%s\n",
				home.Stats.FreeThrowsPercentage,
				pad.AddSpaces(6),
				"%",
				pad.AddSpaces(6),
				away.Stats.FreeThrowsPercentage,
			),
		)
		homeDefReb, _ := strconv.Atoi(home.Stats.ReboundsDefensive)
		awayDefReb, _ := strconv.Atoi(away.Stats.ReboundsDefensive)
		homeOffReb, _ := strconv.Atoi(home.Stats.ReboundsOffensive)
		awayOffReb, _ := strconv.Atoi(away.Stats.ReboundsOffensive)
		str.WriteString(
			fmt.Sprintf("%d%s%s%s%d\n",
				homeDefReb+homeOffReb,
				pad.AddSpaces(5),
				"Tot Reb",
				pad.AddSpaces(5),
				awayDefReb+awayOffReb,
			),
		)
		str.WriteString(
			fmt.Sprintf("%d%s%s%s%d\n",
				homeOffReb,
				pad.AddSpaces(5),
				"Off Reb",
				pad.AddSpaces(5),
				awayOffReb,
			),
		)
		str.WriteString(
			fmt.Sprintf("%2s%s%s%s%2s\n",
				home.Stats.Assists,
				pad.AddSpaces(5),
				"Assists",
				pad.AddSpaces(5),
				away.Stats.Assists,
			),
		)
		str.WriteString(
			fmt.Sprintf("%2s%s%s%s%2s\n",
				home.Stats.Steals,
				pad.AddSpaces(6),
				"Steals",
				pad.AddSpaces(5),
				away.Stats.Steals,
			),
		)
		str.WriteString(
			fmt.Sprintf("%2s%s%s%s%2s\n",
				home.Stats.Turnovers,
				pad.AddSpaces(4),
				"Turnovers",
				pad.AddSpaces(4),
				away.Stats.Turnovers,
			),
		)

	} else {
		str.WriteString("errrr getting team stats")
	}
	return str.String()
}
