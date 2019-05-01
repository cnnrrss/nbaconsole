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
	HasGameBookPdf        bool      `json:"hasGameBookPdf"`
	IsStartTimeTBD        bool      `json:"isStartTimeTBD"`
	Nugget                Nugget    `json:"nugget"`
	GameDuration          struct {
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
		PVTeam            struct {
			SeedNum        string `json:"seedNum"`
			SeriesWin      string `json:"seriesWin"`
			IsSeriesWinner bool   `json:"isSeriesWinner"`
		} `json:"vTeam"`
		PHTeam struct {
			SeedNum        string `json:"seedNum"`
			SeriesWin      string `json:"seriesWin"`
			IsSeriesWinner bool   `json:"isSeriesWinner"`
		} `json:"hTeam"`
	} `json:"playoffs"`
	Period Period `json:"period"`
	HTeam  Team   `json:"hTeam"`
	VTeam  Team   `json:"vTeam"`
}
type Nugget struct {
	Text string `json:"text"`
}

type Period struct {
	Current       int  `json:"current"`
	Type          int  `json:"type"`
	MaxRegular    int  `json:"maxRegular"`
	IsHalftime    bool `json:"isHalftime"`
	IsEndOfPeriod bool `json:"isEndOfPeriod"`
}
type Team struct {
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
}

func DummyDataScoreboard() DataScoreboard {
	return DataScoreboard{
		Games: []Game{
			Game{
				GameID: "111111",
				Nugget: Nugget{
					Text: "Curry 30pts 10/10 3pts",
				},
				VTeam: Team{
					TriCode: "GSW",
					Score:   "110",
				},
				HTeam: Team{
					TriCode: "LAC",
					Score:   "89",
				},
				Period: Period{
					Current: 4,
				},
				IsGameActivated: true,
				StatusNum:       3,
			},
			Game{
				GameID: "111112",
				Nugget: Nugget{
					Text: "Booker 77 pts",
				},
				VTeam: Team{
					TriCode: "PHX",
					Score:   "77",
				},
				HTeam: Team{
					TriCode: "LAL",
					Score:   "150",
				},
				Period: Period{
					Current: 4,
				},
				IsGameActivated: true,
				StatusNum:       3,
			},
			Game{
				GameID: "222222",
				Nugget: Nugget{
					Text: "Harden 110 pts",
				},
				VTeam: Team{
					TriCode: "HOU",
					Score:   "110",
				},
				HTeam: Team{
					TriCode: "BKN",
					Score:   "121",
				},
				Period: Period{
					Current: 4,
				},
				IsGameActivated: true,
				StatusNum:       3,
			},
			Game{
				GameID: "3331112",
				Nugget: Nugget{
					Text: "Lowry 0/11 3pts",
				},
				VTeam: Team{
					TriCode: "TOR",
					Score:   "58",
				},
				HTeam: Team{
					TriCode: "BOS",
					Score:   "43",
				},
				StatusNum: 2,
				Period: Period{
					Current: 2,
				},
				Clock:           "3:01",
				IsGameActivated: true,
			},
			Game{
				GameID: "1442112",
				Nugget: Nugget{
					Text: "",
				},
				VTeam: Team{
					TriCode: "NYK",
					Score:   "3",
				},
				HTeam: Team{
					TriCode: "CHA",
					Score:   "10",
				},
				StatusNum: 2,
				Period: Period{
					Current: 1,
				},
				Clock:           "11:53",
				IsGameActivated: true,
			},
			Game{
				GameID: "00999912",
				Nugget: Nugget{
					Text: "",
				},
				VTeam: Team{
					TriCode: "NOP",
				},
				HTeam: Team{
					TriCode: "CLE",
				},
				StatusNum: 0,
				Period: Period{
					Current: 0,
				},
				StartTimeEastern: "Wednesday",
				IsGameActivated:  false,
			},
		},
	}
}
