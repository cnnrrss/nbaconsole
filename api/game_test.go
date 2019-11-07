package api

import (
	"testing"
)

func Test__playoffRound(t *testing.T) {
	tests := []struct {
		round    int
		div      string
		expected string
	}{
		{1, "East", "Eastern Conference Quarter Finals"},
		{2, "West", "Western Conference Semi Finals"},
		{3, "", "ern Conference Finals"}, // TODO: fix
		{4, "East", "Finals"},
	}

	for _, tc := range tests {
		actual := playOffRound(tc.round, tc.div)
		if actual != tc.expected {
			t.Errorf("actual [%s] not equal to expected [%s]\n", actual, tc.expected)
		}
	}
}
