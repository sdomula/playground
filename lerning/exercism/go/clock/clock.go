// Package clock provides functionality for calculating
// and displaying the clock time of the day.
package clock

import (
	"fmt"
)

const testVersion = 4

// Clock represents a time without the date. Just hours and minutes of the day.
type Clock struct {
	min, hour int
}

// New creates new Clock with the given hour and minute.
func New(hour, minute int) Clock {
	return norm(Clock{hour: hour, min: minute})
}

// String returns the string representation of the time
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}

// Add adds the given minutes to the clock
func (c Clock) Add(minutes int) Clock {
	c.min += minutes
	c = norm(c)
	return c
}

// normalizes clock
func norm(c Clock) Clock {
	c.hour, c.min = normRaw(c.hour, c.min)
	return c
}

func normRaw(hour, min int) (int, int) {
	if min < 0 {
		n := (-min)/60 + 1
		hour -= n
		min += n * 60
	}
	if min >= 60 {
		n := min / 60
		hour += n
		min -= n * 60
	}
	if hour < 0 {
		n := (-hour)/24 + 1
		hour = hour + n*24
	}
	if hour >= 24 {
		n := hour / 24
		hour -= n * 24
	}
	return hour, min
}
