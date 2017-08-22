// Package house provides funtionality to generate
// the nursery rhyme 'This is the House that Jack Built'.
package house

import (
	"fmt"
	"strings"
)

var testVersion = 1

var things = [...]string{
	"house",
	"malt",
	"rat",
	"cat",
	"dog",
	"cow with the crumpled horn",
	"maiden all forlorn",
	"man all tattered and torn",
	"priest all shaven and shorn",
	"rooster that crowed in the morn",
	"farmer sowing his corn",
	"horse and the hound and the horn",
}

var what = [...]string{
	"Jack built",
	"lay in",
	"ate",
	"killed",
	"worried",
	"tossed",
	"milked",
	"kissed",
	"married",
	"woke",
	"kept",
	"belonged to",
}

// Verse generates the nth verse of the song.
func Verse(n int) string {
	if len(things) != len(what) || n > len(things) || n <= 0 {
		return ""
	}
	s := "This is"
	for i := n - 1; i > 0; i-- {
		s += fmt.Sprintf(" the %s\nthat %s", things[i], what[i])
	}
	s += fmt.Sprintf(" the %s that %s.", things[0], what[0])
	return s
}

// Song generates the whole song.
func Song() string {
	s := []string{}
	for i := range things {
		s = append(s, Verse(i+1))
	}
	return strings.Join(s, "\n\n")
}
