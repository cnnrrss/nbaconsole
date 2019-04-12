package api

import "time"

// type Scoreboard struct {
// 	Parameters struct {
// 		GameDate string `json:GameDate`
// 	} `json:parameters`
// 	ResultSets []LineScore
// }

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

type DataScoreBoard struct {
	Internal struct {
		PubDateTime string `json:"pubDateTime"`
		Xslt        string `json:"xslt"`
		EventName   string `json:"eventName"`
	} `json:"_internal"`
	NumGames int `json:"numGames"`
	Games    []struct {
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
	} `json:"games"`
}
