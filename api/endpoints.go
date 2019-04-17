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

// PlayerProfile issues apiGet request for a player profile
func PlayerProfile(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/playerprofilev2", params)
}

// PlayerInfo issues apiGet request for a player common data profile
func PlayerInfo(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/commonplayerinfo", params)
}

// PlayersInfo issues apiGet request for a all players common data
func PlayersInfo(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/commonallplayers", params)
}

// TeamStats issues apiGet request for a leaguedashteamstats
func TeamStats(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguedashteamstats", params)
}

// TeamSplits issues apiGet request for a teamdashboardbygeneralsplits
func TeamSplits(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/teamdashboardbygeneralsplits", params)
}

// TeamYears issues apiGet request for commonteamyears
func TeamYears(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/commonteamyears", params)
}

// PlayerSplits issues apiGet request for player general splits data
func PlayerSplits(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/playerdashboardbygeneralsplits", params)
}

// Shots issues apiGet request for a shot chart detail
func Shots(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/shotchartdetail", params)
}

// ScoreboardV2 issues apiGet request for a scoreboard v2 data
func ScoreboardV2(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/scoreboardv2", params)
}

// BoxScoreSummary issues apiGet request for a summary of boxscore v2 data
func BoxScoreSummary(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/boxscoresummaryv2", params)
}

// BoxScore issues apiGet request for traditional boxscore v2 data
func BoxScore(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/boxscoretraditionalv2", params)
}

// PlayByPlay issues apiGet request for playbyplay data
func PlayByPlay(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/playbyplay", params)
}

// TeamHistoricalLeaders issues apiGet request historical team leader data
func TeamHistoricalLeaders(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/teamhistoricalleaders", params)
}

// TeamInfoCommon issues apiGet request for common team data
func TeamInfoCommon(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/teaminfocommon", params)
}

// CommonTeamRoster issues apiGet request for a team roster
func CommonTeamRoster(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/commonteamroster", params)
}

// TeamPlayerDashboard issues apiGet request for a  team player profile
func TeamPlayerDashboard(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/teamplayerdashboard", params)
}

// Lineups issues apiGet request for a lineup data
func Lineups(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguedashlineups", params)
}

// PlayerTracking issues apiGet request for a player profile
func PlayerTracking(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguedashptstats", params)
}

// HomepageV2 issues apiGet request for NBA API homepage V2
func HomepageV2(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/homepagev2", params)
}

// AssistTracker issues apiGet request for the assist tracker
func AssistTracker(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/assisttracker", params)
}

// PlayerStats issues apiGet request for a players stats
func PlayerStats(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguedashplayerstats", params)
}

// PlayerClutch issues apiGet request for a players clutch data
func PlayerClutch(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguedashplayerclutch", params)
}

// TeamClutch issues apiGet request for a teams clutch data
func TeamClutch(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguedashteamclutch", params)
}

// PlayerShooting issues apiGet request for a players shot data
func PlayerShooting(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguedashplayerptshot", params)
}

// TeamShooting issues apiGet team shooting stats
func TeamShooting(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguedashteamptshot", params)
}

// LeagueGameLog issues apiGet request for league game log
func LeagueGameLog(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguegamelog", params)
}

// LeagueLeaders issues apiGet request for a league leader data
func LeagueLeaders(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leagueLeaders", params)
}

// LeagueStandings issues apiGet request for current league standings
func LeagueStandings(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguestandings", params)
}

// PlayerHustleLeaders issues apiGet request for leage hustle stat leaders
func PlayerHustleLeaders(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguehustlestatsplayerleaders", params)
}

// TeamHustleLeaders issues apiGet request for team leaders in hustle stats
func TeamHustleLeaders(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguehustlestatsteamleaders", params)
}

// PlayerHustle issues apiGet request for player hustle stats data
func PlayerHustle(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguehustlestatsplayer", params)
}

// TeamHustle issues apiGet request for a teams' hustle stats data
func TeamHustle(params map[string]string) (resp *http.Response, err error) {
	return apiGet(Origin+"/stats/leaguehustlestatsteam", params)
}
