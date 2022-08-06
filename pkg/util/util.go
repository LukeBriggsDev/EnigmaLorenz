package util

// negMod is a helper function to have a mod n produce a positive result even if a is negative.
func NegMod(a int, n int) int {
	return (a%n + n) % n
}

// validChars takes a string and returns whether the string only contains ASCII characters in the range A-Z.
func ValidChars(text string) bool {
	for _, chr := range text {
		if chr < 'A' || chr > 'Z' {
			return false
		}
	}
	return true
}
