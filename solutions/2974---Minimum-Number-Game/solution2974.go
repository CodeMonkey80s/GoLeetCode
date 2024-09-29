package solution2974

import (
	"slices"
)

// ============================================================================
// 2974. Minimum Number Game
// URL: https://leetcode.com/problems/minimum-number-game/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2974---Minimum-Number-Game
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	BenchmarkNumberGame
	BenchmarkNumberGame-24    	147935254	         8.680 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func numberGame(nums []int) []int {
	slices.Sort(nums)
	for i := 0; i < len(nums); i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}

	return nums
}
