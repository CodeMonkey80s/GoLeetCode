package solution525

import (
	"fmt"
)

func findMaxLength(nums []int) int {

	prefix := make([]int, len(nums))

	if nums[0] == 1 {
		prefix[0] = 1
	}
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] == 1 {
			prefix[i] = prefix[i-1] + 1
		} else {
			prefix[i] = prefix[i-1] - 1
		}
	}

	fmt.Printf("sum: %v\n", prefix)

	return 0
}

func findMaxLengthV2(nums []int) int {

	prefixSumL := make([]int, len(nums))
	prefixSumR := make([]int, len(nums))

	maxOnes := 0
	maxZeros := 0

	if nums[0] == 0 {
		prefixSumL[0] = 1
	}
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] == 0 {
			prefixSumL[i] = prefixSumL[i-1] + 1
		} else {
			prefixSumL[i] = -1
		}

		maxZeros = max(maxZeros, prefixSumL[i])
	}

	if nums[len(nums)-1] == 1 {
		prefixSumR[len(nums)-1] = 1
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == 1 {
			prefixSumR[i] = prefixSumR[i+1] + 1
		} else {
			prefixSumR[i] = -1
		}

		maxOnes = max(maxOnes, prefixSumR[i])
	}

	// prefixSumL = slices.Insert(prefixSumL, len(nums), 0)
	// prefixSumR = slices.Insert(prefixSumR, 0, 0)
	fmt.Printf("sumL: %v\n", prefixSumL)
	fmt.Printf("sumR: %v\n", prefixSumR)

	// for i := 0; i < len(nums); i++ {
	// 	if prefixSumL[i] == prefixSumR[i] {
	// 		m = max(m, prefixSumL[i]+prefixSumR[i])
	// 	}
	// }

	return maxOnes + maxZeros

	// maxOnes := 0
	// maxZeros := 0
	//
	// countOnes := 0
	// countZeros := 0
	//
	// last := nums[0]
	// for i := 0; i <= len(nums)-1; i++ {
	// 	switch {
	// 	case nums[i] == 0:
	// 		countZeros++
	// 		maxZeros = max(maxZeros, countZeros)
	// 		if last == 1 {
	// 			maxZeros = 0
	// 		}
	// 	case nums[i] == 1:
	// 		countOnes++
	// 		maxOnes = max(maxOnes, countOnes)
	// 		if last == 0 {
	// 			maxOnes = 0
	// 		}
	// 	}
	//
	// }
	//
	// fmt.Printf("max: 0: %d, 1: %d\n", maxZeros, maxOnes)
	// fmt.Printf("cur: 0: %d, 1: %d\n", countZeros, countOnes)
	//
	// return countZeros + countOnes
}
