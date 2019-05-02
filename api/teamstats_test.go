package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestTeamStats(t *testing.T) {
	f, _ := ioutil.ReadFile("./examples/data_api_game_response.json")
	gm := &GameBoxScore{}
	json.Unmarshal(f, gm)
	fmt.Println(gm.TeamStats())
}
