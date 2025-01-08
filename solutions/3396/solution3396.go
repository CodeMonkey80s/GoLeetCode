package solution3396

// ============================================================================
// 3396. Minimum Number of Operations to Make Elements in Array Distinct
// URL: https://leetcode.com/problems/minimum-number-of-operations-to-make-elements-in-array-distinct/
// ============================================================================

func minimumOperations(nums []int) int {

	dist := make([]int, 101)

	for i := len(nums) - 1; i >= 0; i-- {
		dist[nums[i]]++
		if dist[nums[i]] >= 2 {
			return i/3 + 1
		}
	}

	return 0
}
