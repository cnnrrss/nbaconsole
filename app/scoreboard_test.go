package app

import (
	"fmt"
	"testing"

	"github.com/cnnrrss/nbaconsole/api"
)

func Test__formatGame(t *testing.T) {
	sb := api.DummyDataScoreboard()
	for _, gm := range sb.Games {
		fmt.Printf(formatGame(gm))
	}
}
