package solution3467

import "slices"

func transformArray(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[i] = 0
			continue
		}

		nums[i] = 1
	}

	slices.Sort(nums)

	return nums
}
