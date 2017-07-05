// Package hamming provides functionality to calculate the hamming distance.
package hamming

import "fmt"

const testVersion = 6

// Distance calculates the hamming distance between two given strings.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("sequences have different length: len(a)=%d, len(b)=%d", len(a), len(b))
	}
	var dist int
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist, nil
}
