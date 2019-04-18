package api

import (
	"fmt"
	"net/http"
)

func GetDataScoreBoard(params map[string]string) (resp *http.Response, err error) {
	url := fmt.Sprintf("http://data.nba.net/10s/prod/v2/%s/scoreboard.json", params["gameDate"])
	return apiGet(url, nil)
}
