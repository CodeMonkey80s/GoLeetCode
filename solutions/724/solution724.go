package solution724

// ============================================================================
// 724. Find Pivot Index
// URL: https://leetcode.com/problems/find-pivot-index/
// ============================================================================

func pivotIndexV2(nums []int) int {
	var total int
	var leftSum int

	for _, num := range nums {
		total += num
	}

	for i, num := range nums {
		if leftSum == total-leftSum-num {
			return i
		}
		leftSum += num
	}

	return -1
}

func pivotIndex(nums []int) int {

	if len(nums) <= 1 {
		return 0
	}

	prefixSumL := make([]int, len(nums))
	prefixSumL[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSumL[i] = prefixSumL[i-1] + nums[i]
	}

	prefixSumR := make([]int, len(nums))
	prefixSumR[len(prefixSumR)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		prefixSumR[i] = prefixSumR[i+1] + nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 && prefixSumR[i+1] == 0 {
			return i
		}
		if i == len(nums)-1 && prefixSumL[i-1] == 0 {
			return i
		}
		if i > 0 && i < len(nums)-1 && prefixSumL[i-1] == prefixSumR[i+1] {
			return i
		}
	}

	return -1
}
