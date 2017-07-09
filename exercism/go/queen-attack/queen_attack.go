package queenattack

import (
	"fmt"
	"regexp"
)

const testVersion = 2

func CanQueenAttack(a, b string) (bool, error) {
	if a == b || !isValid(a) || !isValid(b) {
		return false, fmt.Errorf("invalid parameter a=%s, b%s", a, b)
	}
	ar, ac := int(a[0]), int(a[1])
	br, bc := int(b[0]), int(b[1])
	// check same row or column
	if ar == br || ac == bc {
		return true, nil
	}
	if abs(ar-br) == abs(ac-bc) {
		return true, nil
	}
	// can not attack
	return false, nil
}

var algNotation = regexp.MustCompile("^([a-h])([1-8])$")

func isValid(notation string) bool {
	m := algNotation.FindStringSubmatch(notation)
	return len(m) == 3 && m[0] == notation
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
