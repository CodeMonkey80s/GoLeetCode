package solution1816

import (
	"strings"
)

// ============================================================================
// 1816. Truncate Sentence
// URL: https://leetcode.com/problems/truncate-sentence/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1816---Truncate-Sentence
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_truncateSentence-24    	14010805	        94.54 ns/op	     104 B/op	       2 allocs/op
	PASS

*/

func truncateSentence(s string, k int) string {
	words := strings.Fields(s)
	if len(words) >= k {
		ans := strings.Join(words[0:k], " ")
		return ans
	}
	return s
}
