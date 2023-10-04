package solution1805

import (
	"strings"
)

// ============================================================================
// 1805. Number of Differenct Integers in a String
// URL: https://leetcode.com/problems/number-of-different-integers-in-a-string/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_numDifferent-24    	10727433	       120.4 ns/op	      80 B/op	       2 allocs/op
	PASS

*/

func numDifferentIntegers(word string) int {
	st := []byte(word)
	for i, ch := range st {
		if '0' <= ch && ch <= '9' {
			continue
		}
		st[i] = ' '
	}
	words := strings.Fields(string(st))
	m := make(map[string]int)
	for _, w := range words {
		val := strings.TrimLeft(w, "0")
		_, ok := m[val]
		if !ok {
			m[val] = 1
		}
	}
	return len(m)
}
