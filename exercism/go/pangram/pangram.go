// Package pangram provides functionality to check for pangrams.
//
// A pangram (Greek: παν γράμμα, pan gramma,
// "every letter") is a sentence using every letter of the alphabet at least once.
// The best known English pangram is:
// > The quick brown fox jumps over the lazy dog.
package pangram

import (
	"strings"
)

const testVersion = 1

// IsPangram determines if a sentence is a pangram.
//
// The alphabet used consists of ASCII letters `a` to `z`, inclusive, and is case
// insensitive. It is assumed that the input will not contain non-ASCII symbols.
func IsPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}
	alphabet := make(map[rune]int, 26)
	for _, r := range strings.ToLower(sentence) {
		if r >= 'a' && r <= 'z' {
			alphabet[r]++
		}
	}
	return len(alphabet) == 26
}
