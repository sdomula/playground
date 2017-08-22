// Package twelve provides functionality to sing 'The Twelve Days of Christmas'.
package twelve

import (
	"fmt"
	"strings"
)

const testVersion = 1

// Song generates the whole song.
func Song() string {
	s := []string{}
	for i := 0; i < len(verses); i++ {
		s = append(s, Verse(i+1))
	}
	return strings.Join(s, "\n") + "\n"
}

// Verse generates the nth verse of the song.
func Verse(n int) string {
	s := fmt.Sprintf("On the %s day of Christmas my true love gave to me", verses[n-1].day)
	if n == 1 {
		return fmt.Sprintf("%s, %s.", s, verses[0].what)
	}
	for i := n - 1; i > 0; i-- {
		s = fmt.Sprintf("%s, %s", s, verses[i].what)
	}
	return fmt.Sprintf("%s, and %s.", s, verses[0].what)
}

var verses = []struct {
	day, what string
}{
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}
