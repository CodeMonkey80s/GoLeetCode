package solution1929

// ============================================================================
// 1929. Concatenation of Array
// URL: https://leetcode.com/problems/concatenation-of-array/
// ============================================================================

func getConcatenation(nums []int) []int {
	ans := make([]int, len(nums)*2)
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[i]
		ans[len(nums)+i] = nums[i]
	}
	return ans

}
