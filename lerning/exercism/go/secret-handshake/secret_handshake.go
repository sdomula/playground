// Package secret provides functionality to compute a secret handshake
package secret

var testVersion = 2

// Handshake can be used to translate a number into the secret handshake
// ```
// 1 = wink
// 10 = double blink
// 100 = close your eyes
// 1000 = jump
//
// 10000 = Reverse the order of the operations in the secret handshake.
// ```
func Handshake(n uint) []string {
	s := []string{}
	if n&1 == 1 {
		s = append(s, "wink")
	}
	if n&2 == 2 {
		s = append(s, "double blink")
	}
	if n&4 == 4 {
		s = append(s, "close your eyes")
	}
	if n&8 == 8 {
		s = append(s, "jump")
	}
	if n&16 == 16 {
		s = reverse(s)
	}
	return s
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
