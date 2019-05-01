package app

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/cnnrrss/nbaconsole/api"
)

func Test__setGames(t *testing.T) {
	console := &NBAConsole{
		gamesList: &Box{},
	}

	scoreboard := api.DummyDataScoreboard()
	console.setGames(scoreboard)
	buf := bytes.NewBuffer([]byte{})
	console.DrawScoreBoard(buf, 120)
	fmt.Println(buf.String())
}
