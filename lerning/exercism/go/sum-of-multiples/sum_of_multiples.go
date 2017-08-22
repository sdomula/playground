package summultiples

const testVersion = 2

func SumMultiples(limit int, divisors ...int) int {
	for _, d := range divisors {
		if d <= 0 {
			return 0
		}
	}
	sum := 0
	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if i%d == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
