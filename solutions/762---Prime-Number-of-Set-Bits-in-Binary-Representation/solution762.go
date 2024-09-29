package solution762

import (
	"math/bits"
)

// ============================================================================
// 762. Prime Number of Set Bits in Binary Representation
// URL: https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/762---Prime-Number-of-Set-Bits-in-Binary-Representation
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countPrimeSetBits
	Benchmark_countPrimeSetBits-24    	328723824	         9.411 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func countPrimeSetBits(left int, right int) int {

	isPrime := func(n int) bool {
		if n <= 1 {
			return false
		}
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	n := 0
	for i := left; i <= right; i++ {
		c := bits.OnesCount(uint(i))
		if isPrime(c) {
			n++
		}
	}

	return n
}
