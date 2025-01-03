package solution1796

import (
	"slices"
)

// ============================================================================
// 1796. Second Largest Digit in a String
// URL: https://leetcode.com/problems/second-largest-digit-in-a-string/
// ============================================================================

/*

	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1796---Second-Largest-Digit-in-a-String
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_secondHighest
	Benchmark_secondHighest-24    	13378276	        82.89 ns/op	      56 B/op	       3 allocs/op
	PASS

*/

func secondHighest(s string) int {
	val := 0
	st := make([]int, 0)
outer:
	for _, ru := range s {
		ch := byte(ru)
		if '0' <= ch && ch <= '9' {
			val = int(ch - 48)
			for _, v := range st {
				if v == val {
					continue outer
				}
			}
			st = append(st, int(ch-48))
		}
	}
	slices.Reverse(st)
	if len(st) > 1 {
		return st[1]
	}
	return -1
}
