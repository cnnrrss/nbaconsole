package api

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestTeamStats(t *testing.T) {
	expected := `WAS     STATS     DET
43/93    FGs      4/6
46.2      %      41.5
8/29     3pts   13/27
27.6      %      48.1
1/25     FTs    18/23
84.0      %      78.3
11     Tot Reb     42
3      Off Reb     13
27     Assists     21
9       Steals      4
1     Turnovers    19`

	f, _ := ioutil.ReadFile("./examples/data_api_game_response.json")
	gm := &GameBoxScore{}
	json.Unmarshal(f, gm)
	if gm.TeamStats() != expected {
		t.Fatalf("got\n%s\nexpected\n%s", gm.TeamStats(), expected)
	}
}
