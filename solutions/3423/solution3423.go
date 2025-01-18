package solution3423

import "math"

// ============================================================================
// 3423. Maximum Difference Between Adjacent Elements in a Circular Array
// URL: https://leetcode.com/problems/maximum-difference-between-adjacent-elements-in-a-circular-array/
// ============================================================================

func maxAdjacentDistance(nums []int) int {

	m := math.MinInt

	for i := 0; i < len(nums); i++ {

		idx1 := i % len(nums)
		idx2 := (i + 1) % len(nums)
		m = max(m, int(math.Abs(float64(nums[idx2]-nums[idx1]))))
	}

	return m
}
