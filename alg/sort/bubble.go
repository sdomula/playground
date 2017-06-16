package sort

func bubble(s []int) []int {
	for x := 0; x < len(s); x++ {
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
			}
		}
	}
	return s
}
