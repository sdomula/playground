package letter

import "sync"

const testVersion = 1

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s []string) FreqMap {
	cfreq := struct {
		mu sync.Mutex
		m  FreqMap
	}{m: FreqMap{}}

	var wg sync.WaitGroup
	for _, str := range s {
		wg.Add(1)
		go func(s string) {
			f := Frequency(s)
			cfreq.mu.Lock()
			for k, v := range f {
				cfreq.m[k] += v
			}
			cfreq.mu.Unlock()
			wg.Done()
		}(str)
	}
	wg.Wait()

	return cfreq.m
}
