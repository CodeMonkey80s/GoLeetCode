package solution228

import (
	"fmt"
	"strconv"
)

// ============================================================================
// 228. Summary Ranges
// URL: https://leetcode.com/problems/summary-ranges/
// ============================================================================

func summaryRanges(nums []int) []string {
	ans := []string{}
	if len(nums) == 0 {
		return ans
	}
	a := 0
	b := 1
	sa := ""
	sb := ""
	for i := 0; i < len(nums); i++ {
		v := strconv.Itoa(nums[i])
		if a == 0 {
			sa = v
		}
		if nums[a] == nums[b]-1 {
			sb = v
			b++
		}
		a++
		fmt.Printf("%d %d, %q %q\n", a, b, sa, sb)
	}

	fmt.Printf("ANS: %s\n", ans)
	return ans
}
