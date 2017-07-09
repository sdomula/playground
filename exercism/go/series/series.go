package series

const testVersion = 2

func All(n int, s string) []string {
	var sub []string
	for i := 0; n <= len(s); i, n = i+1, n+1 {
		sub = append(sub, s[i:n])
	}
	return sub
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return ""
	}
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
