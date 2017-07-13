package etl

import "strings"

const testVersion = 1

func Transform(old map[int][]string) map[string]int {
	new := map[string]int{}
	for value, letters := range old {
		for _, c := range letters {
			new[strings.ToLower(c)] = value
		}
	}
	return new
}
