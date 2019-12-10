package app

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/cnnrrss/nbaconsole/api"
	"github.com/cnnrrss/nbaconsole/common/pad"
)

func (nba *NBAConsole) toggleGameBoxScore() error {
	currentGameID := nba.getSelectedGameID()
	if nba.selectedGameID != currentGameID {
		nba.selectedGameID = currentGameID
	}

	go nba.getBoxScore()

	return nil
}

func (nba *NBAConsole) getBoxScore() error {
	params := genericParams(nba.date)
	if nba.selectedGameID == "" {
		nba.debuglog("error nba selected game nil")
	}

	gameBoxScore := &api.GameBoxScore{}

	// if final game not cached. // TODO: should cache all games.
	if nba.selectedGameScore == nil ||
		nba.selectedGameScore.GameID() != nba.selectedGameID ||
		nba.selectedGameScore.SportsContent.Game.PeriodTime.GameStatus == "3" {

		resp, err := api.GetDataGameBoxScore(params, nba.selectedGameID)
		if err != nil {
			return fmt.Errorf("Error with boxscore request %v", err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Error reading request body %v", err)
		}

		if err := json.Unmarshal(body, &gameBoxScore); err != nil {
			nba.debuglog(fmt.Sprintf("err unmarshalling %v\n", err.Error()))
		}
	}

	nba.selectedGameScore = gameBoxScore

	nba.update(func() {
		nba.boxScore.Clear()
		nba.drawBoxScore(
			nba.boxScore,
			nba.selectedGameScore,
			0,
			false,
		)
	})

	return nil
}

func (nba *NBAConsole) drawBoxScore(output io.Writer, bs *api.GameBoxScore, width int, drawLeaders bool) {
	var str strings.Builder

	str.WriteString(fmt.Sprintf("%-25s%-4s%-4s%-4s%-5s%s\n", "Team", "1", "2", "3", "4", "T"))
	str.WriteString(fmt.Sprintf("%s\n", pad.AddString(str.Len(), "-")))

	hLine, hTotal := homeLineScores(bs)
	str.WriteString(
		fmt.Sprintf("%-24s%-5s%4d\n",
			api.NBATeamDictionary[bs.SportsContent.Game.Home.Abbreviation],
			hLine,
			hTotal,
		),
	)
	vLine, vTotal := visitorLineScores(bs)
	str.WriteString(
		fmt.Sprintf("%-24s%-5s%4d\n",
			api.NBATeamDictionary[bs.SportsContent.Game.Visitor.Abbreviation],
			vLine,
			vTotal,
		),
	)

	fmt.Fprintln(output, str.String())
	if drawLeaders {
		fmt.Fprintln(output, bs.BoxScoreLeaders())
	}

	highlightView(nba.scoreboard)
}

func homeLineScores(bs *api.GameBoxScore) (string, int) {
	var lineScore strings.Builder
	var total int
	curPeriod := len(bs.SportsContent.Game.Home.Linescores.Period)

	for i := 0; i < curPeriod || i <= 3; i++ {
		if i >= curPeriod {
			lineScore.WriteString("  - ")
		} else {
			q, _ := strconv.Atoi(bs.SportsContent.Game.Home.Linescores.Period[i].Score)
			total += q
			lineScore.WriteString(
				fmt.Sprintf("%2d  ", q),
			)
		}
	}
	return lineScore.String(), total
}

func visitorLineScores(bs *api.GameBoxScore) (string, int) {
	var lineScore strings.Builder
	var total int
	curPeriod := len(bs.SportsContent.Game.Visitor.Linescores.Period)

	for i := 0; i < curPeriod || i <= 3; i++ {
		if i >= curPeriod {
			lineScore.WriteString("  - ")
		} else {
			q, _ := strconv.Atoi(bs.SportsContent.Game.Visitor.Linescores.Period[i].Score)
			total += q
			lineScore.WriteString(
				fmt.Sprintf("%2d  ", q),
			)
		}
	}
	return lineScore.String(), total
}
