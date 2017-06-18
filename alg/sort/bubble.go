package sort

func bubble(s []int) []int {
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
