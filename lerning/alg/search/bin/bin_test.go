package bin

import "testing"

func TestSearchOne(t *testing.T) {
	i := Search([]int{1}, 1)
	e := 0
	if i != e {
		t.Errorf("%v != %v", i, e)
	}
}

func TestSearchNotInList(t *testing.T) {
	i := Search([]int{1, 2, 3, 41, 60}, 42)
	e := -1
	if i != e {
		t.Errorf("%v != %v", i, e)
	}
}

func TestSearchInLeft(t *testing.T) {
	i := Search([]int{1, 2, 3}, 1)
	e := 0
	if i != e {
		t.Errorf("%v != %v", i, e)
	}
}

func TestSearchInRight(t *testing.T) {
	i := Search([]int{1, 2, 3}, 3)
	e := 2
	if i != e {
		t.Errorf("%v != %v", i, e)
	}
}

func TestSearchInLargeRight(t *testing.T) {
	i := Search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 2)
	e := 1
	if i != e {
		t.Errorf("%v != %v", i, e)
	}
}
func TestSearchInLargeLeft(t *testing.T) {
	i := Search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 11)
	e := 10
	if i != e {
		t.Errorf("%v != %v", i, e)
	}
}
