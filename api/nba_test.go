package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TodaysGames(t *testing.T) {
	s := NewClient()
	resp, err := s.TodaysGames("")
	if err != nil {
		fmt.Printf("Test failed because: %v", err)
		t.Fail()
	}
	fmt.Println(resp)
	assert.NotEqual(t, resp, "")
}

func Test_FetchNBAToday(t *testing.T) {
	resp, err := FetchNBAToday()
	fmt.Println(resp)
	if err != nil {
		fmt.Printf("Test failed because: %v", err)
		t.Fail()
	}
}
