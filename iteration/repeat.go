package iteration

// Repeat takes string and repeatCount number and returns repeatedString
func Repeat(a string, repeatCount int) string {
	out := ""
	for i := 0; i < repeatCount; i++ {
		out += a
	}
	return out
}
