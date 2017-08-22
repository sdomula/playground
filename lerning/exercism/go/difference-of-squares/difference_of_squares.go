// Package diffsquares provides functionality to
// ind the difference between the square of the sum and the sum of the squares of the first N natural numbers.
package diffsquares

var testVersion = 1

// SquareOfSums returns the sum of the first n natural numbers.
// E.g, `n=10`: `(1 + 2 + ... + 10)² = 55² = 3025`.
func SquareOfSums(n int) int {
	s := 0
	for i := 1; i <= n; i++ {
		s += i
	}
	return s * s
}

// SumOfSquares returns the sum of the squares of the first n natural numbers.
// E.g., given `n=10`: `1² + 2² + ... + 10² = 385`.
func SumOfSquares(n int) int {
	s := 0
	for i := 1; i <= n; i++ {
		s += i * i
	}
	return s
}

// Difference calculates the difference of the squre of sum and the sum of squares
// of the firset n natural numbers.
func Difference(n int) int {
	var s, sq int
	for i := 1; i <= n; i++ {
		s += i
		sq += i * i
	}
	return (s * s) - sq
}
