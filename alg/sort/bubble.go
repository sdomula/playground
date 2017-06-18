package sort

func bubble(s []int) []int {
	for i := 0; i < len(s); i++ {
		if !bubbleLoop(s, i) {
			break
		}
	}
	return s
}

func bubbleLoop(s []int, prev int) bool {
	swapped := false
	for i := 0; i+1 < len(s)-prev; i++ {
		if s[i] > s[i+1] {
			s[i], s[i+1] = s[i+1], s[i]
			swapped = true
		}
	}
	return swapped
}
