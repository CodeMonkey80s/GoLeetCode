package solution228

import (
	"strconv"
)

// ============================================================================
// 228. Summary Ranges
// URL: https://leetcode.com/problems/summary-ranges/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/___unsolved/228---Summary-Ranges
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_summaryRanges
	Benchmark_summaryRanges-24    	 6932410	       154.7 ns/op	     120 B/op	       5 allocs/op
	PASS
*/

func summaryRanges(nums []int) []string {

	var ans []string

	if len(nums) == 0 {
		return ans
	}

	var a, b int

	for b < len(nums) {
		if a != b {
			if b == len(nums)-1 || nums[b] != nums[b+1]-1 {
				n1 := strconv.Itoa(nums[a])
				n2 := strconv.Itoa(nums[b])
				ans = append(ans, n1+"->"+n2)
				a = b + 1
			}
		} else {
			if b == len(nums)-1 || nums[b] < nums[b+1]-1 {
				n1 := strconv.Itoa(nums[b])
				ans = append(ans, n1)
				a = b + 1
			}
		}

		b++
	}

	return ans
}
