package solution2595

import (
	"strconv"
)

// ============================================================================
// 2595. Number of Even and Odd Bits
// URL: https://leetcode.com/problems/number-of-even-and-odd-bits/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2595---Number-of-Even-and-Odd-Bits
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_evenOddBit-24                   	71339138	        16.52 ns/op	       0 B/op	       0 allocs/op
	Benchmark_evenOddBit_first_approach-24    	41538038	        32.68 ns/op	      21 B/op	       2 allocs/op
	PASS

*/

func evenOddBit(n int) []int {
	ans := []int{0, 0}
	for i := 32; i >= 0; i-- {
		if val := n & (1 << i); val > 0 {
			if i%2 == 0 {
				ans[0]++
			} else {
				ans[1]++
			}
		}
	}
	return ans
}

func evenOddBit_first_approach(n int) []int {
	ans := make([]int, 2)
	s := strconv.FormatInt(int64(n), 2)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '1' {
			if len(s)%2 == 0 {
				if i%2 == 0 {
					ans[1]++
				} else {
					ans[0]++
				}
			} else {
				if i%2 == 0 {
					ans[0]++
				} else {
					ans[1]++
				}
			}
		}
	}
	return ans
}
