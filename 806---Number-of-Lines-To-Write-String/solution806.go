package solution806

// ============================================================================
// 806. Number of Lines To Write String
// URL: https://leetcode.com/problems/number-of-lines-to-write-string
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_numberOfLines-24    	70308742	        16.61 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func numberOfLines(widths []int, s string) []int {
	lines := 1
	pixels := 0
	for i := range s {
		ch := s[i]
		w := widths[int(ch-97)]
		if 1 <= pixels+w && pixels+w <= 100 {
			pixels += w
		} else {
			pixels = w
			lines++
		}
	}
	return []int{lines, pixels}
}
