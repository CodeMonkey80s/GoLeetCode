package solution303

// ============================================================================
// 303. Range Sum Query - Immutable
// URL: https://leetcode.com/problems/range-sum-query-immutable/
// ============================================================================

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}

	n := NumArray{
		nums: nums,
	}

	return n
}

func (n *NumArray) SumRange(left int, right int) int {
	if left-1 < 0 {
		return n.nums[right]
	}

	return n.nums[right] - n.nums[left-1]
}
