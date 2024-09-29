package solution963

import (
	"math/bits"
)

// ============================================================================
// 693. Binary Number with Alternating Bits
// URL: https://leetcode.com/problems/binary-number-with-alternating-bits/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/693---Binary-Number-with-Alternating-Bits
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_hasAlternatingBits
	Benchmark_hasAlternatingBits-24    	186342273	         6.426 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func hasAlternatingBits(n int) bool {
	m := bits.LeadingZeros(uint(n))
	for i := 0; i < 64-m; i++ {
		b1 := n & (1 << i)
		b2 := n & (1 << (i + 1))
		if b1 == 0 && b2 == 0 || b1 != 0 && b2 != 0 {
			return false
		}
	}

	return true
}
