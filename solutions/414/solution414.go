package solution414

import (
	"slices"
)

func thirdMax(nums []int) int {
	slices.Sort(nums)
	nums = slices.Compact(nums)
	if len(nums) >= 3 {
		return nums[len(nums)-3]
	}
	return nums[len(nums)-1]
}
