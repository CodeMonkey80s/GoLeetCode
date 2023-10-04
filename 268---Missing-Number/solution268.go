package solution268

// ============================================================================
// 268. Missing Number
// URL: https://leetcode.com/problems/missing-number/
// ============================================================================

func missingNumber(nums []int) int {
	s1, s2 := 0, 0
	size := len(nums)
	for i := 0; i <= size; i++ {
		s1 += i
		if i < size {
			s2 += nums[i]
		}
	}
	return s1 - s2
}
