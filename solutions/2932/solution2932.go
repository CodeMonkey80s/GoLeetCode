package solution2932

import (
	"math"
)

// ============================================================================
// 2932. Maximum Strong Pair XOR I
// URL: https://leetcode.com/problems/maximum-strong-pair-xor-i/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2932---Maximum-Strong-Pair-XOR-I
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	BenchmarkMaximumStrongPairXor
	BenchmarkMaximumStrongPairXor-24    	25561774	        46.92 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func maximumStrongPairXor(nums []int) int {
	var v, m, maxXor int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			v = int(math.Abs(float64(nums[i]) - float64(nums[j])))
			if v <= min(nums[i], nums[j]) {
				maxXor = nums[i] ^ nums[j]
				m = max(m, maxXor)
			}

		}
	}

	return m
}
