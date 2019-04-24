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
}
