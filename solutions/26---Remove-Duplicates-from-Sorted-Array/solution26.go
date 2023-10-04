package solution26

import (
	"slices"
)

// ============================================================================
// 26. Remove Duplicates from Sorted Array
// URL: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// ============================================================================

func removeDuplicates_stdlib(nums []int) int {
	nums = slices.Compact(nums)
	return len(nums)
}

func removeDuplicates(nums []int) int {
	m := make(map[int]int)
	s := make([]int, 0)
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n]++
			s = append(s, n)
		}
	}
	copy(nums, s)
	l2 := len(s)
	nums = nums[0:l2:l2]
	return len(nums)
}
