package solution2974

import "sort"

// ============================================================================
// 2974. Minimum Number Game
// URL: https://leetcode.com/problems/minimum-number-game/
// ============================================================================

func numberGame(nums []int) []int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}

	return nums
}
