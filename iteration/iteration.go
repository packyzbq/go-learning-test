package iteration

import (
	"strings"
)

// Repeat is
func Repeat(s string) (r string) {
	r = strings.Repeat(s, 10)
	return
}

// RepeatN method
func RepeatN(s string, n int) (r string) {
	for i := 0; i < n; i++ {
		r += s
	}
	return
}
