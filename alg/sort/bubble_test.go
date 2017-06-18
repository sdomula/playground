package sort

import "testing"

func eq(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func assertEq(t *testing.T, a, b []int) {
	if !eq(a, b) {
		t.Errorf("Expected %v, got %v", b, a)
	}
}

func TestBubble(t *testing.T) {
	a := bubble([]int{2, 1})
	b := []int{1, 2}
	assertEq(t, a, b)

	a = bubble([]int{3, 2, 1})
	b = []int{1, 2, 3}
	assertEq(t, a, b)

	a = bubble([]int{1, 2, 3})
	b = []int{1, 2, 3}
	assertEq(t, a, b)

	a = bubble([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assertEq(t, a, b)

	a = bubble([]int{3234972, 2, -1, -234, 0, 1})
	b = []int{-234, -1, 0, 1, 2, 3234972}
	assertEq(t, a, b)
}

func TestNil(t *testing.T) {
	a := bubble([]int{})
	b := []int{}
	assertEq(t, a, b)

	a = bubble([]int{})
	var b2 []int
	assertEq(t, a, b2)
}
