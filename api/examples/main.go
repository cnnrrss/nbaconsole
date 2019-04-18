package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	api "github.com/connorvanderhook/nbaconsole/api"
)

// Game is a struct of the game
type Game struct {
	GameCode      string
	HomeTeamCity  string
	HomeScore     float64
	AwayTeamCity  string
	AwayScore     float64
	Quarter       string
	TimeRemaining string
}

var (
	DefaultClient *Client = &Client{
		requester: &DefaultRequester,
	}
)

// Client is a wrapper around the requester struct
type Client struct {
	requester *Requester
}

const (
	NBAStatsDomain = "stats.nba.com"
)

// DefaultRequester is the default Requester using default values for the
// endpoints.
var DefaultRequester = Requester{
	Domain:     NBAStatsDomain,
	PathPrefix: "stats",
}

// Requester performs requests to the stats.nba.com server's endpoints.
type Requester struct {
	Domain     string
	PathPrefix string
	client     http.Client
}

func main() {
	params := map[string]string{
		"gameDate": "20190418",
	}

	resp, err := api.GetDataScoreBoard(params)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(body))
	sb := unMarshalDataScoreboardToStruct(body)
	formatGames(sb)
}

func formatGames(sb api.DataScoreboard) {
	for _, game := range sb.Games {
		if game.Playoffs.RoundNum != "" {
			// if round num = 3 "Conference Finals"
			fmt.Println(game.Playoffs.SeriesSummaryText)
		}

		fmt.Println("Teams |", " Q1 |", "Q2 |", "Q3 |", "Q4 |", "Total")
		fmt.Printf("%-5s | ", game.HTeam.TriCode)
		for _, q := range game.HTeam.Linescore {
			fmt.Printf(" %3v|", q.Score)
		}
		fmt.Printf(" %s \n", game.HTeam.Score)
		fmt.Printf("%-5s | ", game.VTeam.TriCode)
		for _, q := range game.VTeam.Linescore {
			fmt.Printf(" %3v|", q.Score)
		}
		fmt.Printf(" %s \n", game.VTeam.Score)
	}
}

func unMarshalDataScoreboardToStruct(body []byte) api.DataScoreboard {
	results := &api.DataScoreboard{}
	json.Unmarshal(body, &results)
	return *results
}

func toScoreboardStruct(body []byte) {
	results := &api.Scoreboard{}
	json.Unmarshal(body, &results)
	games := []Game{}
	for i, result := range results.ResultSets {

		if result.Name == "GameHeader" {
			for _, row := range result.RowSet {
				game := Game{
					GameCode: row[2].(string),
				}
				games = append(games, game)
			}
		}
		if result.Name == "LineScore" {
			fmt.Println(i, result.Headers[5], result.Headers[8], result.Headers[9], result.Headers[10], result.Headers[11])
			for _, row := range result.RowSet {
				fmt.Println(i, row[5], row[8], row[9], row[10], row[11])
			}
		}
		for _, game := range games {
			fmt.Println(game.GameCode)
		}
	}
}

func toMap(data map[string][]map[string]interface{}) {
	for _, result := range data["resultSets"] {
		// s := result["resultSets"].(map[string]interface{})
		// fmt.Println(result)
		if result["name"] == "LineScore" {
			// fmt.Println(k, result["headers"])
			for _, header := range result["headers"].([]interface{}) {
				fmt.Printf("|%s", header)
			}
			fmt.Printf("|\n")
			for _, row := range result["rowSet"].([]interface{}) {
				fmt.Printf("%v\n", row)
			}
		}
	}
}
