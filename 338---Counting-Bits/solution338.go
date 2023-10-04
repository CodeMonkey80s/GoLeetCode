package solution338

import (
	"math/bits"
)

// ============================================================================
// 338. Counting Bits
// URL: https://leetcode.com/problems/counting-bits/
// ============================================================================

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

func countBits(n int) []int {
	l := n + 1
	s := make([]int, 0, l)
	v := 0
	for i := 0; i <= n; i++ {
		v = bits.OnesCount(uint(i))
		s = append(s, v)
	}
	return s
}
