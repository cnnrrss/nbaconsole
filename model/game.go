package model

var leaderTypes = []string{"Pts", "Reb", "Ast", "Stl", "Blk"}

type Game struct {
	HomeTeam      string
	HomeScore     int
	AwayTeam      string
	AwayScore     int
	Quarter       string
	TimeRemaining string
	Favorite      bool
	Leaders       []GameLeaders
}

type GameLeaders struct {
	PlayerName, Type string
	Value            int
}

type Players struct {
	Players []Player
}

type Player struct {
	Name, Team string
	Number     int
}
