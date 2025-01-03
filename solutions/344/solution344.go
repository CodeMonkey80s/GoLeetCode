package solution344

// ============================================================================
// 344. Reverse String
// URL: https://leetcode.com/problems/reverse-string/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/344---Reverse-String
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_reverse-24    	1000000000	         1.157 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func reverseString(s []byte) {
	size := len(s)
	for i := 0; i < size/2; i++ {
		s[i], s[size-1-i] = s[size-1-i], s[i]
	}
}
