package solution3370

import "math"

// ============================================================================
// 3370. Smallest Number With All Set Bits
// URL: https://leetcode.com/problems/smallest-number-with-all-set-bits/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3370---Smallest-Number-With-All-Set-Bits
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_isBalanced
	Benchmark_isBalanced-24    	53495989	        26.21 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func smallestNumber(n int) int {
	var i int
	var v int
	for {
		v = int(math.Pow(float64(2), float64(i)))
		if v > n {
			break
		}
		i++
	}

	return v - 1
}
