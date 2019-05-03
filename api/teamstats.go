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
		str.WriteString(fmt.Sprintf("%s%s%s%s%s\n", home.Abbreviation, pad.AddString(5, " "), "STATS", pad.AddString(5, " "), away.Abbreviation))
		str.WriteString(
			fmt.Sprintf("%s/%s%s%s%s%s/%s\n",
				home.Stats.FieldGoalsMade,
				home.Stats.FieldGoalsAttempted,
				pad.AddString(4, " "),
				"FGs",
				pad.AddString(4, " "),
				away.Stats.FieldGoalsMade,
				away.Stats.FieldGoalsAttempted,
			),
		)
		str.WriteString(
			fmt.Sprintf("%s%s%s%s%s\n",
				home.Stats.FieldGoalsPercentage,
				pad.AddString(6, " "),
				"%",
				pad.AddString(6, " "),
				away.Stats.FieldGoalsPercentage,
			),
		)
		str.WriteString(
			fmt.Sprintf("%s/%s%s%s%s%s/%s\n",
				home.Stats.ThreePointersMade,
				home.Stats.ThreePointersAttempted,
				pad.AddString(4, " "),
				"3pts",
				pad.AddString(4, " "),
				away.Stats.ThreePointersMade,
				away.Stats.ThreePointersAttempted,
			),
		)
		str.WriteString(
			fmt.Sprintf("%s%s%s%s%s\n",
				home.Stats.ThreePointersPercentage,
				pad.AddString(6, " "),
				"%",
				pad.AddString(6, " "),
				away.Stats.ThreePointersPercentage,
			),
		)
		str.WriteString(
			fmt.Sprintf("%s/%s%s%s%s%s/%s\n",
				home.Stats.FreeThrowsMade,
				home.Stats.FreeThrowsAttempted,
				pad.AddString(4, " "),
				"FTs",
				pad.AddString(4, " "),
				away.Stats.FreeThrowsMade,
				away.Stats.FreeThrowsAttempted,
			),
		)
		str.WriteString(
			fmt.Sprintf("%s%s%s%s%s\n",
				home.Stats.FreeThrowsPercentage,
				pad.AddString(6, " "),
				"%",
				pad.AddString(6, " "),
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
				pad.AddString(5, " "),
				"Tot Reb",
				pad.AddString(5, " "),
				awayDefReb+awayOffReb,
			),
		)
		str.WriteString(
			fmt.Sprintf("%d%s%s%s%d\n",
				homeOffReb,
				pad.AddString(5, " "),
				"Off Reb",
				pad.AddString(5, " "),
				awayOffReb,
			),
		)
		str.WriteString(
			fmt.Sprintf("%2s%s%s%s%2s\n",
				home.Stats.Assists,
				pad.AddString(5, " "),
				"Assists",
				pad.AddString(5, " "),
				away.Stats.Assists,
			),
		)
		str.WriteString(
			fmt.Sprintf("%2s%s%s%s%2s\n",
				home.Stats.Steals,
				pad.AddString(6, " "),
				"Steals",
				pad.AddString(5, " "),
				away.Stats.Steals,
			),
		)
		str.WriteString(
			fmt.Sprintf("%2s%s%s%s%2s\n",
				home.Stats.Turnovers,
				pad.AddString(4, " "),
				"Turnovers",
				pad.AddString(4, " "),
				away.Stats.Turnovers,
			),
		)

	} else {
		str.WriteString("errrr getting team stats")
	}
	return str.String()
}
