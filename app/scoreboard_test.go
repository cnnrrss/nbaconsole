package app

import (
	"fmt"
	"testing"

	"github.com/cnnrrss/nbaconsole/api"
)

func Test__formatGames(t *testing.T) {
	sb := api.DummyDataScoreboard()
	for _, gm := range sb.Games {
		fmt.Printf(formatGames(gm))
	}
}
