package solution392

// ============================================================================
// 392. Is Subsequence
// URL: https://leetcode.com/problems/is-subsequence/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_isSubsequence-24    	182476050	         6.631 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func isSubsequence(s string, t string) bool {
	c := 0
	b := 0
outer:
	for i1 := 0; i1 < len(s); i1++ {
		ru1 := s[i1]
		for i2 := b; i2 < len(t); i2++ {
			ru2 := t[i2]
			if ru1 == ru2 && i1 < len(s)-1 && i2 == len(t)-1 {
				return false
			}
			if ru1 == ru2 {
				if i2+len(s) > i1+len(t) {
					return false
				}
				c++
				b = i2 + 1
				continue outer
			}
		}
	}
	return c >= len(s)
}
