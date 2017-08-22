// Package gigasecond provides functionality to work with
// a gigasecond.
package gigasecond

import "time"

const testVersion = 4

// AddGigasecond adds a gigasecond to the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
