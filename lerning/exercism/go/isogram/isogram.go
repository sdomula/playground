package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

// IsIsogram checks if the given string is an isogram
// (contains no letter twice).
func IsIsogram(s string) bool {
	letters := make(map[rune]bool, len(s))
	for _, r := range strings.ToLower(s) {
		switch {
		case !unicode.IsLetter(r):
			continue
		case letters[r]:
			return false
		default:
			letters[r] = true
		}
	}
	return true
}
