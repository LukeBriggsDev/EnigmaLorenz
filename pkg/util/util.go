package util

import "fmt"

// NegMod is a helper function to have a mod n produce a positive result even if a is negative.
func NegMod(a int, n int) int {
	return (a%n + n) % n
}

func strContains(needle rune, haystack string) bool {
	found := false
	for _, char := range haystack {
		if char == needle {
			found = true
		}
	}

	return found
}

// ValidChars takes a string and returns whether the string only contains ASCII characters in the range A-Z.
// allowNum also allows numbers 0-9 to be valid
func ValidChars(text string, lorenzMode bool) bool {
	special := "(%/-'+\"\\:&)_ £,.0123456789^|=#$*@-<>[]?;{}±"
	for _, chr := range text {
		if chr < 'A' || chr > 'Z' {
			if !lorenzMode {
				return false
			}
			if !strContains(chr, special) {
				fmt.Printf("%c\n", chr)
				return false
			}
		}
	}
	return true
}
