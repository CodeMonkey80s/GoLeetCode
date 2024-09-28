package solution961

import (
	"slices"
)

// ============================================================================
// 961. N-Repeated Element in Size 2N Array
// URL: https://leetcode.com/problems/n-repeated-element-in-size-2n-array/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/961---N-Repeated-Element-in-Size-2N-Array
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_repeatedNTimesV2
	Benchmark_repeatedNTimesV2-24    	25389780	        58.44 ns/op	       0 B/op	       0 allocs/op
	Benchmark_repeatedNTimesV1
	Benchmark_repeatedNTimesV1-24    	100000000	        10.88 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func repeatedNTimesV2(nums []int) int {
	freq := make(map[int]int, len(nums))
	for _, n := range nums {
		if _, ok := freq[n]; ok {
			return n
		}
		freq[n]++
	}

	return 0
}

func repeatedNTimesV1(nums []int) int {
	slices.Sort(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}

	return 0
}
