package solution35

// ============================================================================
// 35. Search Insert Position
// https://leetcode.com/problems/search-insert-position/
// ============================================================================

func searchInsert(nums []int, target int) int {
	n := len(nums)
	i := n / 2
	switch {
	case nums[0] > target:
		return 0
	case nums[i] == target:
		return i
	case nums[i] < target:
		for {
			i++
			if i >= n {
				return i
			}
			if nums[i] == target {
				return i
			}
			if nums[i] > target {
				return i
			}
		}
	case nums[i] > target:
		for {
			i--
			if i < 0 {
				return i
			}
			if nums[i] == target {
				return i
			}
			if nums[i] < target {
				return i + 1
			}
		}
	}
	return 0
}
