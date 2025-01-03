package solution1876

// ============================================================================
// 1876. Substrings of Size Three with Distinct Characters
// URL: https://leetcode.com/problems/substrings-of-size-three-with-distinct-characters/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1876--Substrings-of-Size-Three-with-Distinct-Characters
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countGoodSubstrings-24    	788855241	         1.516 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func countGoodSubstrings(s string) int {
	ans := 0
	for i := 0; i < len(s)-2; i++ {
		a := s[i]
		b := s[i+1]
		c := s[i+2]
		if a == b || b == c || c == a {
			continue
		}
		ans++
	}
	return ans
}
