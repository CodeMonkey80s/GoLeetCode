package solution1051

import (
	"slices"
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
	Benchmark_heightChecker-24    	38655662	        30.87 ns/op	      48 B/op	       1 allocs/op
	PASS

*/

func heightChecker(heights []int) int {
	ans := 0
	sorted := make([]int, len(heights))
	copy(sorted, heights)
	slices.Sort(sorted)
	for i, v := range sorted {
		if v != heights[i] {
			ans++
		}
	}
	return ans
}
