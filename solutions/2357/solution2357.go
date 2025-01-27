package solution2357

import (
	"slices"
)

func minimumOperations(nums []int) int {

	slices.Sort(nums)

	var total int
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	var ops int
	for i := 0; i < len(nums); i++ {

		val := nums[i]
		if val == 0 {
			continue
		}

		ops++
		for j := i; j < len(nums); j++ {
			nums[j] = nums[j] - val
		}

		if nums[len(nums)-1] == 0 {
			break
		}
	}

	return ops
}
