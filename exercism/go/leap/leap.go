// Package leap provides a function to check
// the given year.
package leap

const testVersion = 3

// IsLeapYear returns true if the given year is a year that
// contains an extra day (February 29).
func IsLeapYear(y int) bool {
	if y%4 != 0 {
		return false
	}
	if y%100 == 0 && y%400 != 0 {
		return false
	}
	return true
}
