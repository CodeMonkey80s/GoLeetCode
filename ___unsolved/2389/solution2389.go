package solution2389

// ============================================================================
// 2389. Longest Subsequence With Limited Sum
// URL: https://leetcode.com/problems/longest-subsequence-with-limited-sum/
// ============================================================================

import (
	"slices"
)

func answerQueries(nums []int, queries []int) []int {

	output := make([]int, len(queries))

	slices.Sort(nums)

	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}

loop:
	for idx, q := range queries {
		c := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= q {
				c++
			} else {
				output[idx] = c
				continue loop
			}
		}
		output[idx] = c
	}

	return output
}
