package api

import (
	"fmt"
	"net/http"
)

// GetDataScoreBoard issues an apiGet request on the data API scoreboard endpoint
func GetDataScoreBoard(params map[string]string) (resp *http.Response, err error) {
	url := fmt.Sprintf("http://data.nba.net/10s/prod/v2/%s/scoreboard.json", params["gameDate"])
	return apiGet(url, nil)
}

// GetDataGameBoxScore issues an apiGet request on the data API boscore endpoint
func GetDataGameBoxScore(params map[string]string, gameID string) (resp *http.Response, err error) {
	url := fmt.Sprintf("http://data.nba.net/data/5s/json/cms/noseason/game/%s/%s/boxscore.json", params["gameDate"], gameID)
	return apiGet(url, nil)
}
