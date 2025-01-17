package solution2903

import "math"

// ============================================================================
// 2903. Find Indices With Index and Value Difference I
// URL: https://leetcode.com/problems/find-indices-with-index-and-value-difference-i/
// ============================================================================

func findIndices(nums []int, indexDifference int, valueDifference int) []int {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if int(math.Abs(float64(i-j))) >= indexDifference && int(math.Abs(float64(nums[i]-nums[j]))) >= valueDifference {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}
