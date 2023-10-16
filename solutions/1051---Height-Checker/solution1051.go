package solution1051

import (
	"sort"
)

// ============================================================================
// 1051. Height Checker
// URL: https://leetcode.com/problems/height-checker/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1051---Height-Checker
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_heightChecker-24    	17535759	        63.04 ns/op	      72 B/op	       2 allocs/op
	PASS

*/

func heightChecker(heights []int) int {
	ans := 0
	sorted := make([]int, len(heights))
	copy(sorted, heights)
	sort.Ints(sorted)
	for i, v := range sorted {
		if v != heights[i] {
			ans++
		}
	}
	return ans
}
