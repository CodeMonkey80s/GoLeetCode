package soluton3300

import "slices"

func minElement(nums []int) int {

	sumOfDigits := func(num int) int {
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		return sum
	}

	for i, num := range nums {
		sum := sumOfDigits(num)
		nums[i] = sum
	}

	return slices.Min(nums)
}
