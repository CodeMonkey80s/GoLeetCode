package solution283

import (
	"sort"
)

// ============================================================================
// 283. Move Zeroes
// URL: https://leetcode.com/problems/move-zeroes/
// ============================================================================

func moveZeroes(nums []int) {
	sort.Slice(nums, func(i int, j int) bool {
		return nums[j] == 0
	})
}
