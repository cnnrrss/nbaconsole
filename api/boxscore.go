package api

import (
	"fmt"
	"strings"
)

func (bs *GameBoxScore) PointsLeaders() string {
	var str strings.Builder

	home, away := bs.SportsContent.Game.Home, bs.SportsContent.Game.Visitor
	str.WriteString(
		fmt.Sprintf("%s %s %s\n",
			home.Leaders.Points.Leader[0].FirstName,
			home.Leaders.Points.Leader[0].LastName,
			home.Leaders.Points.StatValue,
		))
	str.WriteString(
		fmt.Sprintf("%s %s %s\n",
			away.Leaders.Points.Leader[0].FirstName,
			away.Leaders.Points.Leader[0].LastName,
			away.Leaders.Points.StatValue,
		))
	return str.String()
}

func (bs *GameBoxScore) AssistsLeaders() {
	return
}

func (bs *GameBoxScore) ReboundsLeaders() {
	return
}

type GameBoxScore struct {
	SportsContent struct {
		Game struct {
			Arena      string `json:"arena"`
			Attendance string `json:"attendance"`
			City       string `json:"city"`
			Country    string `json:"country"`
			Date       string `json:"date"`
			GameURL    string `json:"game_url"`
			Home       struct {
				Leaders struct {
					Assists struct {
						PlayerCount string `json:"PlayerCount"`
						StatValue   string `json:"StatValue"`
						Leader      []struct {
							FirstName  string `json:"FirstName"`
							LastName   string `json:"LastName"`
							PersonID   string `json:"PersonID"`
							PlayerCode string `json:"PlayerCode"`
						} `json:"leader"`
					} `json:"Assists"`
					Points struct {
						PlayerCount string `json:"PlayerCount"`
						StatValue   string `json:"StatValue"`
						Leader      []struct {
							FirstName  string `json:"FirstName"`
							LastName   string `json:"LastName"`
							PersonID   string `json:"PersonID"`
							PlayerCode string `json:"PlayerCode"`
						} `json:"leader"`
					} `json:"Points"`
					Rebounds struct {
						PlayerCount string `json:"PlayerCount"`
						StatValue   string `json:"StatValue"`
						Leader      []struct {
							FirstName  string `json:"FirstName"`
							LastName   string `json:"LastName"`
							PersonID   string `json:"PersonID"`
							PlayerCode string `json:"PlayerCode"`
						} `json:"leader"`
					} `json:"Rebounds"`
				} `json:"Leaders"`
				Abbreviation string `json:"abbreviation"`
				City         string `json:"city"`
				ID           string `json:"id"`
				Linescores   struct {
					Period []struct {
						PeriodName  string `json:"period_name"`
						PeriodValue string `json:"period_value"`
						Score       string `json:"score"`
					} `json:"period"`
				} `json:"linescores"`
				Nickname string `json:"nickname"`
				Players  struct {
					Player []struct {
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
					} `json:"player"`
				} `json:"players"`
				Score string `json:"score"`
				Stats struct {
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
				} `json:"stats"`
				TeamCode string `json:"team_code"`
				TeamKey  string `json:"team_key"`
				URLName  string `json:"url_name"`
			} `json:"home"`
			HomeStartDate string `json:"home_start_date"`
			HomeStartTime string `json:"home_start_time"`
			ID            string `json:"id"`
			PeriodTime    struct {
				GameClock    string `json:"game_clock"`
				GameStatus   string `json:"game_status"`
				PeriodName   string `json:"period_name"`
				PeriodStatus string `json:"period_status"`
				PeriodValue  string `json:"period_value"`
				TotalPeriods string `json:"total_periods"`
			} `json:"period_time"`
			PreviewAvailable string `json:"previewAvailable"`
			RecapAvailable   string `json:"recapAvailable"`
			SeasonID         string `json:"season_id"`
			State            string `json:"state"`
			Time             string `json:"time"`
			TntOt            string `json:"tnt_ot"`
			Visitor          struct {
				Leaders struct {
					Assists struct {
						PlayerCount string `json:"PlayerCount"`
						StatValue   string `json:"StatValue"`
						Leader      []struct {
							FirstName  string `json:"FirstName"`
							LastName   string `json:"LastName"`
							PersonID   string `json:"PersonID"`
							PlayerCode string `json:"PlayerCode"`
						} `json:"leader"`
					} `json:"Assists"`
					Points struct {
						PlayerCount string `json:"PlayerCount"`
						StatValue   string `json:"StatValue"`
						Leader      []struct {
							FirstName  string `json:"FirstName"`
							LastName   string `json:"LastName"`
							PersonID   string `json:"PersonID"`
							PlayerCode string `json:"PlayerCode"`
						} `json:"leader"`
					} `json:"Points"`
					Rebounds struct {
						PlayerCount string `json:"PlayerCount"`
						StatValue   string `json:"StatValue"`
						Leader      []struct {
							FirstName  string `json:"FirstName"`
							LastName   string `json:"LastName"`
							PersonID   string `json:"PersonID"`
							PlayerCode string `json:"PlayerCode"`
						} `json:"leader"`
					} `json:"Rebounds"`
				} `json:"Leaders"`
				Abbreviation string `json:"abbreviation"`
				City         string `json:"city"`
				ID           string `json:"id"`
				Linescores   struct {
					Period []struct {
						PeriodName  string `json:"period_name"`
						PeriodValue string `json:"period_value"`
						Score       string `json:"score"`
					} `json:"period"`
				} `json:"linescores"`
				Nickname string `json:"nickname"`
				Players  struct {
					Player []struct {
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
					} `json:"player"`
				} `json:"players"`
				Score string `json:"score"`
				Stats struct {
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
				} `json:"stats"`
				TeamCode string `json:"team_code"`
				TeamKey  string `json:"team_key"`
				URLName  string `json:"url_name"`
			} `json:"visitor"`
			VisitorStartDate string `json:"visitor_start_date"`
			VisitorStartTime string `json:"visitor_start_time"`
		} `json:"game"`
		SportsMeta struct {
			ConsolidatedDomKey string `json:"consolidatedDomKey"`
			DateTime           string `json:"date_time"`
			EndToEndTimeMillis string `json:"end_to_end_time_millis"`
			Next               struct {
				URL string `json:"url"`
			} `json:"next"`
			SeasonMeta struct {
				CalendarDate        string `json:"calendar_date"`
				DisplaySeason       string `json:"display_season"`
				DisplayYear         string `json:"display_year"`
				LeagueID            string `json:"league_id"`
				RosterSeasonYear    string `json:"roster_season_year"`
				ScheduleSeasonYear  string `json:"schedule_season_year"`
				SeasonID            string `json:"season_id"`
				SeasonStage         string `json:"season_stage"`
				SeasonYear          string `json:"season_year"`
				StandingsSeasonYear string `json:"standings_season_year"`
				StatsSeasonID       string `json:"stats_season_id"`
				StatsSeasonStage    string `json:"stats_season_stage"`
				StatsSeasonYear     string `json:"stats_season_year"`
			} `json:"season_meta"`
		} `json:"sports_meta"`
	} `json:"sports_content"`
}
