package app

import (
	"github.com/jroimartin/gocui"
	log "github.com/sirupsen/logrus"
)

// update passes in a function to update an NBAConsole view
func (nba *NBAConsole) update(f func()) {
	if nba.g == nil {
		log.Fatal("gocui is not initialized")
	}

	nba.g.Update(func(g *gocui.Gui) error {
		f()
		return nil
	})
}

// PadCenter splits the current length in half and pads each side of the string
func PadCenter(str string, length int) string {
	return Pad(" ", length/2-len(str)) + str + Pad(" ", length/2-len(str))
}

// PadLeft pads the string on the left side up to length-str
func PadLeft(str string, length int) string {
	return Pad(" ", length-len(str)) + str
}

// PadRight pads the string on the right side up to length-str
func PadRight(str string, length int) string {
	return str + Pad(" ", length-len(str))
}

// Pad takes a string and prints the string n times
func Pad(str string, n int) string {
	var out string
	for i := 0; i < n; i++ {
		out += str
	}
	return out
}
