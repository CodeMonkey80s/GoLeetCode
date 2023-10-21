package solution2506

// ============================================================================
// 2506. Count Pairs of Similar Strings
// URL: https://leetcode.com/problems/count-pairs-of-similar-strings/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2506---Count-Pairs-of-Similar-Strings
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_similarPairs-24    	24235304	        48.49 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func similarPairs(words []string) int {
	ans := 0
	m1 := 0
	m2 := 0
	setBits := func(word string) int {
		m := 0
		for _, ch := range word {
			m |= 1 << int(ch-'a')
		}
		return m
	}
	for i := 0; i < len(words); i++ {
		m1 = setBits(words[i])
		for j := i; j < len(words); j++ {
			if i < j {
				m2 = setBits(words[j])
				if m1 == m2 {
					ans++
				}
			}
		}
	}
	return ans
}
