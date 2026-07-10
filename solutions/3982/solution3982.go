package solution3982

import (
	"math"
	"strconv"
)

func maxDigitRange(nums []int) int {
	ranges := make([]int, len(nums))
	var maxRange int
	for i := 0; i < len(nums); i++ {
		vMax := 0
		vMin := math.MaxInt
		digits := strconv.Itoa(nums[i])
		for _, d := range digits {
			vMax = max(int(d-'0'), vMax)
			vMin = min(int(d-'0'), vMin)
		}
		r := vMax - vMin
		maxRange = max(r, maxRange)
		ranges[i] = r
	}

	var sum int
	for i := 0; i < len(nums); i++ {
		if ranges[i] == maxRange {
			sum += nums[i]
		}
	}

	return sum
}
