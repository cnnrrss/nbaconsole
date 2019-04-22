package app

const version = "0.0.1"

func (nba *NBAConsole) version() string {
	return version
}

// Version returns NBAConsole version
func Version() string {
	return version
}
