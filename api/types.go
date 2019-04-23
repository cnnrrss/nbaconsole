package api

import (
	"time"
)

// Scoreboard is the api response using the stats api
type Scoreboard struct {
	Resource   string `json:"resource"`
	Parameters struct {
		GameDate  string `json:"GameDate"`
		LeagueID  string `json:"LeagueID"`
		DayOffset string `json:"DayOffset"`
	} `json:"parameters"`
	ResultSets []struct {
		Name    string          `json:"name"`
		Headers []string        `json:"headers"`
		RowSet  [][]interface{} `json:"rowSet"`
	} `json:"resultSets"`
}

// DataScoreboard is the api response using the data api
type DataScoreboard struct {
	Internal struct {
		PubDateTime string `json:"pubDateTime"`
		Xslt        string `json:"xslt"`
		EventName   string `json:"eventName"`
	} `json:"_internal"`
	NumGames int    `json:"numGames"`
	Games    []Game `json:"games"`
}

// Game is the game data embedded in the DataScoreboard type
type Game struct {
	SeasonStageID int    `json:"seasonStageId"`
	SeasonYear    string `json:"seasonYear"`
	GameID        string `json:"gameId"`
	Arena         struct {
		Name       string `json:"name"`
		IsDomestic bool   `json:"isDomestic"`
		City       string `json:"city"`
		StateAbbr  string `json:"stateAbbr"`
		Country    string `json:"country"`
	} `json:"arena"`
	IsGameActivated       bool      `json:"isGameActivated"`
	StatusNum             int       `json:"statusNum"`
	ExtendedStatusNum     int       `json:"extendedStatusNum"`
	StartTimeEastern      string    `json:"startTimeEastern"`
	StartTimeUTC          time.Time `json:"startTimeUTC"`
	EndTimeUTC            time.Time `json:"endTimeUTC"`
	StartDateEastern      string    `json:"startDateEastern"`
	Clock                 string    `json:"clock"`
	IsBuzzerBeater        bool      `json:"isBuzzerBeater"`
	IsPreviewArticleAvail bool      `json:"isPreviewArticleAvail"`
	IsRecapArticleAvail   bool      `json:"isRecapArticleAvail"`
	Tickets               struct {
		MobileApp    string `json:"mobileApp"`
		DesktopWeb   string `json:"desktopWeb"`
		MobileWeb    string `json:"mobileWeb"`
		LeagGameInfo string `json:"leagGameInfo"`
		LeagTix      string `json:"leagTix"`
	} `json:"tickets"`
	HasGameBookPdf bool `json:"hasGameBookPdf"`
	IsStartTimeTBD bool `json:"isStartTimeTBD"`
	Nugget         struct {
		Text string `json:"text"`
	} `json:"nugget"`
	Attendance   string `json:"attendance"`
	GameDuration struct {
		Hours   string `json:"hours"`
		Minutes string `json:"minutes"`
	} `json:"gameDuration"`
	Tags     []string `json:"tags"`
	Playoffs struct {
		RoundNum          string `json:"roundNum"`
		ConfName          string `json:"confName"`
		SeriesID          string `json:"seriesId"`
		SeriesSummaryText string `json:"seriesSummaryText"`
		IsSeriesCompleted bool   `json:"isSeriesCompleted"`
		GameNumInSeries   string `json:"gameNumInSeries"`
		IsIfNecessary     bool   `json:"isIfNecessary"`
		VTeam             struct {
			SeedNum        string `json:"seedNum"`
			SeriesWin      string `json:"seriesWin"`
			IsSeriesWinner bool   `json:"isSeriesWinner"`
		} `json:"vTeam"`
		HTeam struct {
			SeedNum        string `json:"seedNum"`
			SeriesWin      string `json:"seriesWin"`
			IsSeriesWinner bool   `json:"isSeriesWinner"`
		} `json:"hTeam"`
	} `json:"playoffs"`
	Period struct {
		Current       int  `json:"current"`
		Type          int  `json:"type"`
		MaxRegular    int  `json:"maxRegular"`
		IsHalftime    bool `json:"isHalftime"`
		IsEndOfPeriod bool `json:"isEndOfPeriod"`
	} `json:"period"`
	VTeam struct {
		TeamID     string `json:"teamId"`
		TriCode    string `json:"triCode"`
		Win        string `json:"win"`
		Loss       string `json:"loss"`
		SeriesWin  string `json:"seriesWin"`
		SeriesLoss string `json:"seriesLoss"`
		Score      string `json:"score"`
		Linescore  []struct {
			Score string `json:"score"`
		} `json:"linescore"`
	} `json:"vTeam"`
	HTeam struct {
		TeamID     string `json:"teamId"`
		TriCode    string `json:"triCode"`
		Win        string `json:"win"`
		Loss       string `json:"loss"`
		SeriesWin  string `json:"seriesWin"`
		SeriesLoss string `json:"seriesLoss"`
		Score      string `json:"score"`
		Linescore  []struct {
			Score string `json:"score"`
		} `json:"linescore"`
	} `json:"hTeam"`
	Watch struct {
		Broadcast struct {
			Broadcasters struct {
				National []struct {
					ShortName string `json:"shortName"`
					LongName  string `json:"longName"`
				} `json:"national"`
				Canadian []struct {
					ShortName string `json:"shortName"`
					LongName  string `json:"longName"`
				} `json:"canadian"`
				VTeam []struct {
					ShortName string `json:"shortName"`
					LongName  string `json:"longName"`
				} `json:"vTeam"`
				HTeam []struct {
					ShortName string `json:"shortName"`
					LongName  string `json:"longName"`
				} `json:"hTeam"`
				SpanishHTeam    []interface{} `json:"spanish_hTeam"`
				SpanishVTeam    []interface{} `json:"spanish_vTeam"`
				SpanishNational []interface{} `json:"spanish_national"`
			} `json:"broadcasters"`
			Video struct {
				RegionalBlackoutCodes string `json:"regionalBlackoutCodes"`
				CanPurchase           bool   `json:"canPurchase"`
				IsLeaguePass          bool   `json:"isLeaguePass"`
				IsNationalBlackout    bool   `json:"isNationalBlackout"`
				IsTNTOT               bool   `json:"isTNTOT"`
				IsVR                  bool   `json:"isVR"`
				TntotIsOnAir          bool   `json:"tntotIsOnAir"`
				IsNextVR              bool   `json:"isNextVR"`
				IsNBAOnTNTVR          bool   `json:"isNBAOnTNTVR"`
				IsMagicLeap           bool   `json:"isMagicLeap"`
				IsOculusVenues        bool   `json:"isOculusVenues"`
				Streams               []struct {
					StreamType            string `json:"streamType"`
					IsOnAir               bool   `json:"isOnAir"`
					DoesArchiveExist      bool   `json:"doesArchiveExist"`
					IsArchiveAvailToWatch bool   `json:"isArchiveAvailToWatch"`
					StreamID              string `json:"streamId"`
					Duration              int    `json:"duration"`
				} `json:"streams"`
				DeepLink []struct {
					Broadcaster         string `json:"broadcaster"`
					RegionalMarketCodes string `json:"regionalMarketCodes"`
					IosApp              string `json:"iosApp"`
					AndroidApp          string `json:"androidApp"`
					DesktopWeb          string `json:"desktopWeb"`
					MobileWeb           string `json:"mobileWeb"`
				} `json:"deepLink"`
			} `json:"video"`
			Audio struct {
				National struct {
					Streams []struct {
						Language string `json:"language"`
						IsOnAir  bool   `json:"isOnAir"`
						StreamID string `json:"streamId"`
					} `json:"streams"`
					Broadcasters []interface{} `json:"broadcasters"`
				} `json:"national"`
				VTeam struct {
					Streams []struct {
						Language string `json:"language"`
						IsOnAir  bool   `json:"isOnAir"`
						StreamID string `json:"streamId"`
					} `json:"streams"`
					Broadcasters []struct {
						ShortName string `json:"shortName"`
						LongName  string `json:"longName"`
					} `json:"broadcasters"`
				} `json:"vTeam"`
				HTeam struct {
					Streams []struct {
						Language string `json:"language"`
						IsOnAir  bool   `json:"isOnAir"`
						StreamID string `json:"streamId"`
					} `json:"streams"`
					Broadcasters []struct {
						ShortName string `json:"shortName"`
						LongName  string `json:"longName"`
					} `json:"broadcasters"`
				} `json:"hTeam"`
			} `json:"audio"`
		} `json:"broadcast"`
	} `json:"watch"`
}

// GameBoxScore is the json response API from data.nba.net for a box score
type GameBoxScore struct {
	SportsContent struct {
		Game struct {
			Arena        string `json:"arena"`
			Attendance   string `json:"attendance"`
			Broadcasters struct {
				Radio struct {
					Broadcaster []struct {
						DisplayName string `json:"display_name"`
						HomeVisitor string `json:"home_visitor"`
						Scope       string `json:"scope"`
					} `json:"broadcaster"`
				} `json:"radio"`
				Tv struct {
					Broadcaster []struct {
						DisplayName string `json:"display_name"`
						HomeVisitor string `json:"home_visitor"`
						Scope       string `json:"scope"`
					} `json:"broadcaster"`
				} `json:"tv"`
			} `json:"broadcasters"`
			City    string `json:"city"`
			Country string `json:"country"`
			Date    string `json:"date"`
			Dl      struct {
				Link []interface{} `json:"link"`
			} `json:"dl"`
			GameURL string `json:"game_url"`
			Home    struct {
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
			Lp            struct {
				CondensedBb string `json:"condensed_bb"`
				Home        struct {
					Audio struct {
						ENG string `json:"ENG"`
						SPA string `json:"SPA"`
					} `json:"audio"`
					Video struct {
						ArchBB string `json:"archBB"`
						Avl    string `json:"avl"`
						OnAir  string `json:"onAir"`
					} `json:"video"`
				} `json:"home"`
				LpVideo string `json:"lp_video"`
				Visitor struct {
					Audio struct {
						ENG string `json:"ENG"`
						SPA string `json:"SPA"`
					} `json:"audio"`
					Video struct {
						ArchBB string `json:"archBB"`
						Avl    string `json:"avl"`
						OnAir  string `json:"onAir"`
					} `json:"video"`
				} `json:"visitor"`
			} `json:"lp"`
			NotebookAvailable string `json:"notebookAvailable"`
			Officials         []struct {
				FirstName    string `json:"first_name"`
				JerseyNumber string `json:"jersey_number"`
				LastName     string `json:"last_name"`
				PersonID     string `json:"person_id"`
			} `json:"officials"`
			PeriodTime struct {
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
			Ticket           struct {
				TicketLink string `json:"ticket_link"`
			} `json:"ticket"`
			Time    string `json:"time"`
			TntOt   string `json:"tnt_ot"`
			Visitor struct {
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
