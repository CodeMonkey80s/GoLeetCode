package solution1078

import (
	"strings"
)

// ============================================================================
// 1078. Occurrences After Bigram
// URL: https://leetcode.com/problems/occurrences-after-bigram/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_find-24    	 8361966	       157.3 ns/op	     208 B/op	       3 allocs/op
	PASS

*/

func findOcurrences(text string, first string, second string) []string {
	out := []string{}
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		if 0 <= i && i < len(words)-1 {
			if i < len(words)-2 && first == words[i] && second == words[i+1] {
				out = append(out, words[i+2])
			}
		}
	}
	return out
}
