package triangle

import "math"

const testVersion = 3

// KindFromSides Determine if a triangle is equilateral, isosceles, or scalene.
//
// An _equilateral_ triangle has all three sides the same length.
// An _isosceles_ triangle has at least two sides the same length.
// A _scalene_ triangle has all sides of different lengths.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case isNaT(a, b, c):
		return NaT
	case isEqu(a, b, c):
		return Equ
	case isSca(a, b, c):
		return Sca
	default:
		return Iso
	}
}

func isNaT(a, b, c float64) bool {
	if (a <= 0 || b <= 0 || c <= 0) || (a+b < c || a+c < b || b+c < a) {
		return true
	}
	if math.IsInf(a+b+c, 0) || math.IsNaN(a+b+c) {
		return true
	}
	return false
}

func isEqu(a, b, c float64) bool {
	return a == b && b == c
}

func isSca(a, b, c float64) bool {
	return a != b && b != c && a != c
}

func isIso(a, b, c float64) bool {
	return a == b || a == c || b == c
}

// Kind is the kind of the triangle
type Kind string

// Possible Kind values
const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)
