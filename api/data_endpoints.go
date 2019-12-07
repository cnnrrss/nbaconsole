package api

import (
	"fmt"
	"net/http"
)

// GetDataScoreBoard issues an apiGet request on the data API scoreboard endpoint
func GetDataScoreBoard(params map[string]string) (resp *http.Response, err error) {
	url := fmt.Sprintf("%s/10s/prod/v2/%s/scoreboard.json",
		_origin,
		params["gameDate"],
	)
	return apiGet(url, nil)
}

// GetDataGameBoxScore issues an apiGet request on the data API boscore endpoint
func GetDataGameBoxScore(params map[string]string, gameID string) (resp *http.Response, err error) {
	url := fmt.Sprintf("%s/data/5s/json/cms/noseason/game/%s/%s/boxscore.json",
		_origin,
		params["gameDate"],
		gameID,
	)
	return apiGet(url, nil)
}
