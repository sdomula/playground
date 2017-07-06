// Package acronym provides functionality to build acronyms.
package acronym

import "strings"

const testVersion = 3

// Abbreviate converts a phrase to its acronym.
// E.g., `Abbreviate(Portable Network Graphics)` does return `"PNG"`.
// Techies love their TLA (Three Letter Acronyms)!
func Abbreviate(s string) string {
	words := strings.FieldsFunc(s, func(c rune) bool {
		return c == ' ' || c == '-'
	})
	a := ""
	for _, word := range words {
		a += strings.ToUpper(string(word[0]))
	}
	return a
}
