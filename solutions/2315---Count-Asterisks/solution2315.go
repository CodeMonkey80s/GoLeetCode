package solution2315

import (
	"strings"
)

// ============================================================================
// 2315.Count Asterisks
// URL: https://leetcode.com/problems/count-asterisks/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2315---Count-Asterisks
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countAsterisks-24        	330401191	         3.682 ns/op	       0 B/op	       0 allocs/op
	Benchmark_countAsterisks_std-24    	23830894	        52.56 ns/op	      80 B/op	       1 allocs/op
	PASS

*/

func countAsterisks(s string) int {
	n := 0
	pair := 0
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '|':
			pair++
		case pair%2 == 0 && s[i] == '*':
			n++
		}
	}
	return n
}

func countAsterisks_std(s string) int {
	words := strings.Split(s, "|")
	n := 0
	for i, word := range words {
		if i%2 == 0 {
			n += strings.Count(word, "*")
		}
	}
	return n
}
