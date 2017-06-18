package bubble

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

func eqStrings(a, b []string) bool {
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

func TestInts(t *testing.T) {
	a := Ints([]int{2, 1})
	b := []int{1, 2}
	if !eq(a, b) {
		t.Errorf("Expected %v, got %v", b, a)
	}

	a = Ints([]int{3, 2, 1})
	b = []int{1, 2, 3}
	if !eq(a, b) {
		t.Errorf("Expected %v, got %v", b, a)
	}

	a = Ints([]int{1, 2, 3})
	b = []int{1, 2, 3}
	if !eq(a, b) {
		t.Errorf("Expected %v, got %v", b, a)
	}

	a = Ints([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !eq(a, b) {
		t.Errorf("Expected %v, got %v", b, a)
	}

	a = Ints([]int{3234972, 2, -1, -234, 0, 1})
	b = []int{-234, -1, 0, 1, 2, 3234972}
	if !eq(a, b) {
		t.Errorf("Expected %v, got %v", b, a)
	}

	a = Ints([]int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12})
	b = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	if !eq(a, b) {
		t.Errorf("Expected %v, got %v", b, a)
	}

}

func TestNil(t *testing.T) {
	a := Ints([]int{})
	b := []int{}
	if !eq(a, b) {
		t.Errorf("Expected %v, got %v", b, a)
	}

	a = Ints([]int{})
	var bs []int
	if !eq(a, b) {
		t.Errorf("Expected %v, got %v", bs, a)
	}
}

func BenchmarkInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ints([]int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12})
	}
}

func TestIntsReverse(t *testing.T) {
	a := IntsReverse([]int{1, 2})
	b := []int{2, 1}
	if !eq(a, b) {
		t.Errorf("Expected %v, got %v", b, a)
	}

	a = IntsReverse([]int{1, 2, 3})
	b = []int{3, 2, 1}
	if !eq(a, b) {
		t.Errorf("Expected %v, got %v", b, a)
	}

	a = IntsReverse([]int{3, 2, 1})
	b = []int{3, 2, 1}
	if !eq(a, b) {
		t.Errorf("Expected %v, got %v", b, a)
	}

	a = IntsReverse([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	b = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	if !eq(a, b) {
		t.Errorf("Expected %v, got %v", b, a)
	}

	a = IntsReverse([]int{3234972, 2, -1, -234, 0, 1})
	b = []int{3234972, 2, 1, 0, -1, -234}
	if !eq(a, b) {
		t.Errorf("Expected %v, got %v", b, a)
	}
}

func BenchmarkIntsReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntsReverse([]int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12})
	}
}

func TestBubbleStrings(t *testing.T) {
	a := Strings([]string{"dog", "cat", "alligator", "cheetah", "rat", "moose", "cow", "mink", "porcupine", "dung beetle", "camel", "steer", "bat", "hamster", "horse", "colt", "bald eagle", "frog", "rooster"})
	b := []string{"alligator", "bald eagle", "bat", "camel", "cat", "cheetah", "colt", "cow", "dog", "dung beetle", "frog", "hamster", "horse", "mink", "moose", "porcupine", "rat", "rooster", "steer"}
	if !eqStrings(a, b) {
		t.Errorf("Expected %v, got %v", b, a)
	}
}
