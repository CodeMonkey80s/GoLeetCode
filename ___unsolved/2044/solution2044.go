package solution2044

import (
	"fmt"
)

func countMaxOrSubsets(nums []int) int {

	for i := 0; i < 1<<len(nums); i++ {

		sub := make([]int, len(nums))
		copy(sub, nums)
		for j := 0; j < len(nums); j++ {
			val := i & (1 << j)
			if val > 0 {
				sub[j] = 0
			}
		}
		fmt.Println(sub)

	}

	return 0
}
