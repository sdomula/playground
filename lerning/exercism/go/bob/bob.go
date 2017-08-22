// Package bob simulates a teenager.
package bob

import (
	"strings"
)

const testVersion = 3

// Hey takes a string and returns Bobs reaction.
//
// Bob is a lackadaisical teenager. In conversation, his responses are very limited.
//
// Bob answers 'Sure.' if you ask him a question.
//
// He answers 'Whoa, chill out!' if you yell at him.
//
// He says 'Fine. Be that way!' if you address him without actually saying
// anything.
//
// He answers 'Whatever.' to anything else.
func Hey(s string) string {
	s = strings.TrimSpace(s)
	switch {
	case len(s) == 0:
		return "Fine. Be that way!"
	case isShouted(s):
		return "Whoa, chill out!"
	case isQuestion(s):
		return "Sure."
	default:
		return "Whatever."
	}
}

func isQuestion(s string) bool {
	return s[len(s)-1] == '?'
}

func isShouted(s string) bool {
	return strings.ToUpper(s) == s && (strings.ToUpper(s) != strings.ToLower(s))
}
