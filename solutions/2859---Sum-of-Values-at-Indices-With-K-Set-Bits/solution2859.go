package solution2859

import "math/bits"

// ============================================================================
// 2859. Sum of Values at Indices With K Set Bits
// URL: https://leetcode.com/problems/sum-of-values-at-indices-with-k-set-bits/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2859---Sum-of-Values-at-Indices-With-K-Set-Bits
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_sumIndices-24    	939698586	         1.244 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func sumIndicesWithKSetBits(nums []int, k int) int {
	ans := 0
	for i, v := range nums {
		c := bits.OnesCount(uint(i))
		if c == k {
			ans += v
		}
	}
	return ans
}
