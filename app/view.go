package app

// Pad takes a string and prints the string n times
func helperPad(str string, n int) string {
	var out string
	for i := 0; i < n; i++ {
		out += str
	}
	return out
}
