package solution561

import "slices"

func arrayPairSum(nums []int) int {
	slices.Sort(nums)

	var sum int
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] < nums[i+1] {
			sum += nums[i]
		} else {
			sum += nums[i+1]
		}
	}

	return sum
}
