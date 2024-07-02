package solution3194

import "sort"

// ============================================================================
// 3194. Minimum Average of Smallest and Largest Elements
// URL: https://leetcode.com/problems/minimum-average-of-smallest-and-largest-elements/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3194---Minimum-Average-of-Smallest-and-Largest-Elements
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_minimumAverage-24    	100000000	        19.10 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func minimumAverage(nums []int) float64 {
	sort.Ints(nums)
	ans := 50.0
	i := 0
	for {
		minVal := nums[0]
		maxVal := nums[len(nums)-1]

		nums = nums[1:]
		nums = nums[:len(nums)-1]

		avg := (float64(minVal) + float64(maxVal)) / 2

		if avg < ans {
			ans = avg
		}

		i++

		if len(nums) == 0 {
			break
		}
	}

	return ans
}
