package solution1844

// ============================================================================
// 1844. Replace All Digits with Characters
// URL: https://leetcode.com/problems/replace-all-digits-with-characters/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1844---Replace-All-Digits-with-Characters
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_replaceDigits-24    	198263569	         6.067 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func replaceDigits(s string) string {
	ans := []byte(s)
	for i := 0; i < len(s)-1; i += 2 {
		a := s[i]
		b := s[i+1] - '0'
		c := a + b
		ans[i] = a
		ans[i+1] = c
	}
	return string(ans)
}
