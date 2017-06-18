package sort

// Bubble sorts and returns the given list
// with bubblesort
func Bubble(s []int) []int {
	var swapped bool
	for x := 0; x < len(s); x++ {
		swapped = false
		for i := 0; i+1 < len(s)-x; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return s
}

// BubbleReverse make a reverse bubble sort.
func BubbleReverse(s []int) []int {
	var swapped bool
	for x := 0; x < len(s); x++ {
		swapped = false
		for i := 0; i+1 < len(s)-x; i++ {
			if s[i] < s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return s
}
