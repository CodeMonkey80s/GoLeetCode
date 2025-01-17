package solution2908

import "math"

func minimumSum(nums []int) int {

	m := math.MaxInt
	f := false
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			for k := j; k < len(nums); k++ {
				if i < j && j < k && nums[i] < nums[j] && nums[k] < nums[j] {
					m = min(nums[i]+nums[j]+nums[k], m)
					f = true
				}
			}
		}
	}

	if f {
		return m
	}

	return -1
}
