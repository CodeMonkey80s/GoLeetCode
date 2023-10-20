package solution1408

import "strings"

// ============================================================================
// 1408. String Matching in an Array
// URL: https://leetcode.com/problems/string-matching-in-an-array/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1408---String-Matching-in-an-Array
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_stringMatching-24    	 7692170	       175.2 ns/op	      48 B/op	       2 allocs/op
	PASS

*/

func stringMatching(words []string) []string {
	freq := make(map[string]int)
	ans := make([]string, 0)
	for i, sub := range words {
		for j, word := range words {
			_, ok := freq[sub]
			if !ok && i != j && strings.Contains(word, sub) {
				freq[sub]++
				ans = append(ans, sub)
			}
		}
	}
	return ans
}
