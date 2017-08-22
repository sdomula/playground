// Package raindrops provides functionality to
// create fancy strings from boring numbers.
package raindrops

import (
	"strconv"
	"strings"
)

const testVersion = 3

const (
	// Pling for 3
	Pling = 3
	// Plang for 5
	Plang = 5
	// Plong for 7
	Plong = 7
)

// Convert a number to a string, the contents of which depend on the number's factors.
//
// - If the number has 3 as a factor, output 'Pling'.
// - If the number has 5 as a factor, output 'Plang'.
// - If the number has 7 as a factor, output 'Plong'.
// - If the number does not have 3, 5, or 7 as a factor,
//   just pass the number's digits straight through.
func Convert(n int) string {
	s := []string{}
	if hasFac(n, Pling) {
		s = append(s, "Pling")
	}
	if hasFac(n, Plang) {
		s = append(s, "Plang")
	}
	if hasFac(n, Plong) {
		s = append(s, "Plong")
	}
	if len(s) == 0 {
		s = append(s, strconv.Itoa(n))
	}
	return strings.Join(s, "")
}

func hasFac(n, f int) bool {
	return n%f == 0
}
