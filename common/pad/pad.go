package pad

import (
	"strings"
)

func times(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return
}

// Left left-pads the string with pad up to len runes
func Left(str string, length int, pad string) string {
	return times(pad, length-len(str)) + str
}

// AddString adds the passed in string _n_ times
func AddString(n int, s string) string {
	var str strings.Builder
	for i := 0; i < n; i++ {
		str.WriteString(s)
	}
	return str.String()
}
