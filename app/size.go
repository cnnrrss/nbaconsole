package nbaconsole

// Size returns window width and height
func (nba *NBAConsole) size() (int, int) {
	return nba.g.Size()
}
