package solution1913

import "sort"

// ============================================================================
// 1408. String Matching in an Array
// URL: https://leetcode.com/problems/string-matching-in-an-array/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1913---Maximum-Product-Difference-Between-Two-Pairs
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_maxProductDifference-24    	52369410	        26.55 ns/op	      24 B/op	       1 allocs/op
	PASS

*/

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	a := nums[0]
	b := nums[1]
	c := nums[len(nums)-1]
	d := nums[len(nums)-2]
	return (c * d) - (a * b)
}
