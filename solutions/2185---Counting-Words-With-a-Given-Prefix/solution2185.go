package solution2185

// ============================================================================
// 2185. Counting Words With a Given Prefix
// URL: https://leetcode.com/problems/counting-words-with-a-given-prefix/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2185---Counting-Words-With-a-Given-Prefix
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_prefixCount-24    	144226126	         8.646 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func prefixCount(words []string, pref string) int {
	ans := 0
	for _, word := range words {
		if len(pref) > len(word) {
			continue
		}
		if pref == word[:len(pref)] {
			ans++
		}
	}
	return ans
}
