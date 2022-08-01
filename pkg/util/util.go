package util

// negMod is a helper function to have a mod n produce a positive result even if a is negative.
func NegMod(a int, n int) int {
	return (a%n + n) % n
}
