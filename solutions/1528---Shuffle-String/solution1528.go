package solution1528

// ============================================================================
// 1528. Shuffle String
// URL: https://leetcode.com/problems/shuffle-string/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_restoreString-24    	76885707	        13.86 ns/op	       8 B/op	       1 allocs/op
	PASS

*/

func restoreString(s string, indices []int) string {
	ans := make([]byte, len(s))
	for i, v := range indices {
		ans[v] = s[i]
	}
	return string(ans)
}
