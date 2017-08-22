package bin

// Search returns the index of i in s, -1 if not found.
func Search(s []int, x int) int {
	start := 0
	end := len(s) - 1

	for start <= end {
		m := start + (end-start)/2
		switch {
		case s[m] == x:
			return m
		case x > s[m]:
			start = m + 1
		case x < s[m]:
			end = m - 1
		}
	}

	return -1
}
