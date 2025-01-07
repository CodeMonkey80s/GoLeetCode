package solution1561

// ============================================================================
// 1561. Maximum Number of Coins You Can Get
// URL: https://leetcode.com/problems/maximum-number-of-coins-you-can-get/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1561
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_maxCoins
	Benchmark_maxCoins-24    	104682013	        18.79 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

import "slices"

func maxCoins(piles []int) int {

	slices.SortFunc(piles, func(a, b int) int {
		return b - a
	})

	coins := 0
	for len(piles) > 0 {
		coins += piles[1]
		piles = piles[2 : len(piles)-1]
	}

	return coins
}
