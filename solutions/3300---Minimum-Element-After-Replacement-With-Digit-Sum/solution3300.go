package soluton3300

import "slices"

// ============================================================================
// 3300. Minimum Element After Replacement With Digit Sum
// URL: https://leetcode.com/problems/minimum-element-after-replacement-with-digit-sum/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3300---Minimum-Element-After-Replacement-With-Digit-Sum
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_minElement
	Benchmark_minElement-24    	100000000	        10.54 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func minElement(nums []int) int {

	sumOfDigits := func(num int) int {
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		return sum
	}

	for i, num := range nums {
		sum := sumOfDigits(num)
		nums[i] = sum
	}

	return slices.Min(nums)
}
