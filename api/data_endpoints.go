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
