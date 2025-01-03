package solution338

import (
	"math/bits"
)

// ============================================================================
// 338. Counting Bits
// URL: https://leetcode.com/problems/counting-bits/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/338---Counting-Bits
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countBits
	Benchmark_countBits-24    	58887084	        19.44 ns/op	      24 B/op	       1 allocs/op
	PASS
*/

/*
	Input: n = 5
	Output: [0,1,1,2,1,2]
	Explanation:
	0 --> 0
	1 --> 1
	2 --> 10
	3 --> 11
	4 --> 100
	5 --> 101
*/

func countBitsV2(n int) []int {
	l := n + 1
	s := make([]int, 0, l)
	v := 0
	for i := 0; i <= n; i++ {
		v = 0
		for j := 0; j <= 63; j++ {
			if val := i & (1 << j); val > 0 {
				v++
			}
		}
		s = append(s, v)
	}
	return s
}

func countBitsV1(n int) []int {
	l := n + 1
	s := make([]int, 0, l)
	v := 0
	for i := 0; i <= n; i++ {
		v = bits.OnesCount(uint(i))
		s = append(s, v)
	}
	return s
}
