package pythagorean

import "math"

const testVersion = 1

type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	ts := []Triplet{}
	for i := min; i <= max-max/3; i++ {
		ipow := i * i
		for j := i + 1; j <= max-1; j++ {
			k := math.Sqrt(float64(ipow + (j * j)))
			if k == math.Trunc(k) {
				ts = append(ts, Triplet{i, j, int(k)})
			}
		}
	}
	return ts
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	tsall := Range(3, p)
	ts := []Triplet{}
	for _, t := range tsall {
		if t[0]+t[1]+t[2] == p {
			ts = append(ts, t)
		}

	}
	return ts
}
