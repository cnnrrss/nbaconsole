package api

import (
	"fmt"
	"strings"

	"github.com/cnnrrss/nbaconsole/common/pad"
)

// BoxScoreLeaders writes the game leaders for the
// main counting stats (points, rebounds, assists)
func (bs *GameBoxScore) BoxScoreLeaders() string {
	str := &strings.Builder{}
	if bs != nil {
		home, visitor := bs.SportsContent.Game.Home, bs.SportsContent.Game.Visitor

		str.WriteString(fmt.Sprintf("%s\n", home.Nickname))
		str.WriteString(fmt.Sprintf("%s\n", pad.AddString(39 /** TODO: get dynamic */, "-")))
		if len(home.Leaders.Points.Leader) > 0 {
			writeStatToString(str, home, "points")
		}

		if len(home.Leaders.Assists.Leader) > 0 {
			writeStatToString(str, home, "assists")
		}

		if len(home.Leaders.Rebounds.Leader) > 0 {
			writeStatToString(str, home, "rebounds")
		}

		str.WriteString(fmt.Sprintf("\n%s\n", visitor.Nickname))
		str.WriteString(fmt.Sprintf("%s\n", pad.AddString(39 /** TODO: get dynamic */, "-")))
		if len(visitor.Leaders.Points.Leader) > 0 {
			writeStatToString(str, visitor, "points")
		}

		if len(visitor.Leaders.Assists.Leader) > 0 {
			writeStatToString(str, visitor, "assists")
		}

		if len(visitor.Leaders.Rebounds.Leader) > 0 {
			writeStatToString(str, visitor, "rebounds")
		}
	} else {
		str.WriteString("errrr getting game leaders")
	}
	return str.String()
}

// GameID returns the game ID string for the GameBoxScore Struct
func (bs *GameBoxScore) GameID() string {
	return bs.SportsContent.Game.ID
}

func writeStatToString(str *strings.Builder, team TeamBoxScore, stat string) {
	switch {
	case stat == "points":
		str.WriteString(
			fmt.Sprintf("%s %s %s %s\n",
				team.Leaders.Points.Leader[0].FirstName,
				team.Leaders.Points.Leader[0].LastName,
				team.Leaders.Points.StatValue,
				stat,
			),
		)
	case stat == "assists":
		str.WriteString(
			fmt.Sprintf("%s %s %s %s\n",
				team.Leaders.Assists.Leader[0].FirstName,
				team.Leaders.Assists.Leader[0].LastName,
				team.Leaders.Assists.StatValue,
				stat,
			),
		)
	case stat == "rebounds":
		str.WriteString(
			fmt.Sprintf("%s %s %s %s\n",
				team.Leaders.Rebounds.Leader[0].FirstName,
				team.Leaders.Rebounds.Leader[0].LastName,
				team.Leaders.Rebounds.StatValue,
				stat,
			),
		)
	}
}

// GameBoxScore is a representation of an nba game boxscore
type GameBoxScore struct {
	SportsContent BoxScoreContent `json:"sports_content"`
}

type BoxScoreContent struct {
	Game BoxScoreGame `json:"game"`
}

type BoxScoreGame struct {
	Home             TeamBoxScore           `json:"home"`
	HomeStartDate    string                 `json:"home_start_date"`
	HomeStartTime    string                 `json:"home_start_time"`
	ID               string                 `json:"id"`
	PeriodTime       BoxScoreGamePeriodTime `json:"period_time"`
	PreviewAvailable string                 `json:"previewAvailable"`
	RecapAvailable   string                 `json:"recapAvailable"`
	SeasonID         string                 `json:"season_id"`
	State            string                 `json:"state"`
	Time             string                 `json:"time"`
	TntOt            string                 `json:"tnt_ot"`
	Visitor          TeamBoxScore           `json:"visitor"`
	VisitorStartDate string                 `json:"visitor_start_date"`
	VisitorStartTime string                 `json:"visitor_start_time"`
}

type BoxScoreGamePeriodTime struct {
	GameClock    string `json:"game_clock"`
	GameStatus   string `json:"game_status"`
	PeriodName   string `json:"period_name"`
	PeriodStatus string `json:"period_status"`
	PeriodValue  string `json:"period_value"`
	TotalPeriods string `json:"total_periods"`
}

// TeamBoxScore details each team of a games' boxscore
type TeamBoxScore struct {
	Leaders      TeamBoxScoreStatsLeader `json:"Leaders"`
	Abbreviation string                  `json:"abbreviation"`
	City         string                  `json:"city"`
	ID           string                  `json:"id"`
	Linescores   TeamBoxScoreLinescores  `json:"linescores"`
	Nickname     string                  `json:"nickname"`
	Players      PlayersBoxStats         `json:"players"`
	Score        string                  `json:"score"`
	Stats        TeamBoxStats            `json:"stats"`
	TeamCode     string                  `json:"team_code"`
	TeamKey      string                  `json:"team_key"`
	URLName      string                  `json:"url_name"`
}

type TeamBoxScoreLinescores struct {
	Period []LineScorePeriod `json:"period"`
}

type TeamBoxScoreStatsLeader struct {
	Assists  TeamBoxScoreStatLeader `json:"Assists"`
	Points   TeamBoxScoreStatLeader `json:"Points"`
	Rebounds TeamBoxScoreStatLeader `json:"Rebounds"`
}

type TeamBoxScoreStatLeader struct {
	PlayerCount string                     `json:"PlayerCount"`
	StatValue   string                     `json:"StatValue"`
	Leader      []BoxScoreLeaderPlayerInfo `json:"leader"`
}

type BoxScoreLeaderPlayerInfo struct {
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	PersonID   string `json:"PersonID"`
	PlayerCode string `json:"PlayerCode"`
}

type LineScorePeriod struct {
	PeriodName  string `json:"period_name"`
	PeriodValue string `json:"period_value"`
	Score       string `json:"score"`
}

type PlayersBoxStats struct {
	Player []PlayerBoxStats `json:"player"`
}

type PlayerBoxStats struct {
	Assists                string `json:"assists"`
	Blocks                 string `json:"blocks"`
	FieldGoalsAttempted    string `json:"field_goals_attempted"`
	FieldGoalsMade         string `json:"field_goals_made"`
	FirstName              string `json:"first_name"`
	Fouls                  string `json:"fouls"`
	FreeThrowsAttempted    string `json:"free_throws_attempted"`
	FreeThrowsMade         string `json:"free_throws_made"`
	JerseyNumber           string `json:"jersey_number"`
	LastName               string `json:"last_name"`
	Minutes                string `json:"minutes"`
	OnCourt                string `json:"on_court"`
	PersonID               string `json:"person_id"`
	PlayerCode             string `json:"player_code"`
	PlusMinus              string `json:"plus_minus"`
	Points                 string `json:"points"`
	PositionFull           string `json:"position_full"`
	PositionShort          string `json:"position_short"`
	ReboundsDefensive      string `json:"rebounds_defensive"`
	ReboundsOffensive      string `json:"rebounds_offensive"`
	Seconds                string `json:"seconds"`
	StartingPosition       string `json:"starting_position"`
	Steals                 string `json:"steals"`
	TeamTurnovers          string `json:"team_turnovers"`
	ThreePointersAttempted string `json:"three_pointers_attempted"`
	ThreePointersMade      string `json:"three_pointers_made"`
	Turnovers              string `json:"turnovers"`
}

type TeamBoxStats struct {
	Assists                 string `json:"assists"`
	Blocks                  string `json:"blocks"`
	FieldGoalsAttempted     string `json:"field_goals_attempted"`
	FieldGoalsMade          string `json:"field_goals_made"`
	FieldGoalsPercentage    string `json:"field_goals_percentage"`
	Fouls                   string `json:"fouls"`
	FreeThrowsAttempted     string `json:"free_throws_attempted"`
	FreeThrowsMade          string `json:"free_throws_made"`
	FreeThrowsPercentage    string `json:"free_throws_percentage"`
	FullTimeoutRemaining    string `json:"full_timeout_remaining"`
	Points                  string `json:"points"`
	ReboundsDefensive       string `json:"rebounds_defensive"`
	ReboundsOffensive       string `json:"rebounds_offensive"`
	ShortTimeoutRemaining   string `json:"short_timeout_remaining"`
	Steals                  string `json:"steals"`
	TeamFouls               string `json:"team_fouls"`
	TeamRebounds            string `json:"team_rebounds"`
	TeamTurnovers           string `json:"team_turnovers"`
	TechnicalFouls          string `json:"technical_fouls"`
	ThreePointersAttempted  string `json:"three_pointers_attempted"`
	ThreePointersMade       string `json:"three_pointers_made"`
	ThreePointersPercentage string `json:"three_pointers_percentage"`
	Turnovers               string `json:"turnovers"`
}
