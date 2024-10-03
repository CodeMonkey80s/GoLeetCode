package solution1863

// ============================================================================
// 1863. Sum of All Subset XOR Totals
// URL: https://leetcode.com/problems/sum-of-all-subset-xor-totals/
// ============================================================================

func subsetXORSum(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans |= nums[i]
	}

	return ans << (len(nums) - 1)
}
