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
	Benchmark_countGoodSubstrings-24    	213954763	         5.636 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func countGoodSubstrings(s string) int {
	ans := 0
	var freq uint32
	for i := 0; i < len(s)-2; i++ {
		freq = 0
		a := s[i] - 'a'
		b := s[i+1] - 'a'
		c := s[i+2] - 'a'
		if freq&(1<<a) > 0 {
			continue
		}
		freq |= 1 << a
		if freq&(1<<b) > 0 {
			continue
		}
		freq |= 1 << b
		if freq&(1<<c) > 0 {
			continue
		}
		ans++
	}
	return ans
}
