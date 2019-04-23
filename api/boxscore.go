package api

import (
	"fmt"
)

func (bs *GameBoxScore) Teams() (struct{}, struct{}) {
	return bs.SportsContent.Game.Home, bs.SportsContent.Game.Visitor
}

func (bs *GameBoxScore) PointsLeaders(home Home, away Visitor) {
	fn := home.Leaders.Points.Leader[0].FirstName
	ln := home.Leaders.Points.Leader[0].LastName
	pts := home.Leaders.Points.StatValue
	fmt.Fprintf("%s %s %s\n", home.Abbreviation, fn, ln, pts)
	fn = away.Leaders.Points.Leader[0].FirstName
	ln = away.Leaders.Points.Leader[0].LastName
	pts = away.Leaders.Points.StatValue
	fmt.Fprintf("%s %s %s\n", away.Abbreviation, fn, ln, pts)
}

func (bs *GameBoxScore) AssistsLeaders() {
	return
}

func (bs *GameBoxScore) ReboundsLeaders() {
	return
}
