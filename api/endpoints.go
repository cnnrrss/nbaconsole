package api

// Potential API endpoint
// "http://data.nba.net/10s/prod/v1/today.json"

import (
	"net/http"
)

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	Referer   = "http://stats.nba.com/scores/"
	Origin    = "http://stats.nba.com"
)

func PlayerProfile(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/playerprofilev2", params)
}

func PlayerInfo(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/commonplayerinfo", params)
}

func PlayersInfo(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/commonallplayers", params)
}

func TeamStats(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguedashteamstats", params)
}

func TeamSplits(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/teamdashboardbygeneralsplits", params)
}

func TeamYears(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/commonteamyears", params)
}

func PlayerSplits(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/playerdashboardbygeneralsplits", params)
}

func Shots(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/shotchartdetail", params)
}

func ScoreboardV2(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/scoreboardv2", params)
}

func BoxScoreSummary(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/boxscoresummaryv2", params)
}

func BoxScore(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/boxscoretraditionalv2", params)
}

func PlayByPlay(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/playbyplay", params)
}

func TeamHistoricalLeaders(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/teamhistoricalleaders", params)
}

func TeamInfoCommon(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/teaminfocommon", params)
}

func CommonTeamRoster(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/commonteamroster", params)
}

func TeamPlayerDashboard(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/teamplayerdashboard", params)
}

func Lineups(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguedashlineups", params)
}

func PlayerTracking(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguedashptstats", params)
}

func HomepageV2(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/homepagev2", params)
}

func AssistTracker(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/assisttracker", params)
}

func PlayerStats(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguedashplayerstats", params)
}

func PlayerClutch(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguedashplayerclutch", params)
}

func TeamClutch(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguedashteamclutch", params)
}

func PlayerShooting(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguedashplayerptshot", params)
}

func TeamShooting(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguedashteamptshot", params)
}

func LeagueGameLog(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguegamelog", params)
}

func LeagueLeaders(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leagueLeaders", params)
}

func LeagueStandings(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguestandings", params)
}

func PlayerHustleLeaders(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguehustlestatsplayerleaders", params)
}

func TeamHustleLeaders(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguehustlestatsteamleaders", params)
}

func PlayerHustle(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguehustlestatsplayer", params)
}

func TeamHustle(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguehustlestatsteam", params)
}
