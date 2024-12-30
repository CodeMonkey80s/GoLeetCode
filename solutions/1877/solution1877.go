package solution1877

// ============================================================================
// 1877. Minimize Maximum Pair Sum in Array
// URL: https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1877
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_minPairSum
	Benchmark_minPairSum-24    	77912378	        21.72 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

import (
	"math"
	"slices"
)

func minPairSum(nums []int) int {

	slices.Sort(nums)

	m := math.MinInt

	a := len(nums)/2 - 1
	b := len(nums) / 2
	for i := 0; i < len(nums); i++ {

		m = max(m, nums[a]+nums[b], nums[a-1]+nums[b+1])

		a--
		b++
		if a <= 0 || b >= len(nums)-1 {
			break
		}
	}

	return m
}
