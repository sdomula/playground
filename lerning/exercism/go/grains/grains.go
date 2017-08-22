package grains

import "fmt"

const testVersion = 1

func Square(s int) (uint64, error) {
	if s <= 0 || s > 64 {
		return 0, fmt.Errorf("square is not valid: %d", s)
	}
	var g uint64 = 1
	for i := 1; i < s; i++ {
		g *= 2
	}
	return g, nil
}

func Total() uint64 {
	return 18446744073709551615
}
