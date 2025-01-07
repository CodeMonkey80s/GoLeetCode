package solution3065

import "slices"

// ============================================================================
// 3065. Minimum Operations to Exceed Threshold Value I
// URL: https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-i/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3065
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_minOperations
	Benchmark_minOperations-24    	120689896	         9.693 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func minOperations(nums []int, k int) int {

	slices.Sort(nums)

	var ops int
	for i := 0; i < len(nums); i++ {
		if nums[i] >= k {
			break
		}

		ops++
	}

	return ops
}
