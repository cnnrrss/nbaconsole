package app

import (
	"fmt"
	"log"
	"os"
	"time"
)

func (nba *NBAConsole) debuglog(msg string) {
	if !nba.debug {
		return
	}

	filename := "/tmp/nbaconsole.log"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	text := fmt.Sprintf("%v %s\n", time.Now().Unix(), msg)
	if _, err = f.WriteString(text); err != nil {
		log.Fatal(err)
	}
}
