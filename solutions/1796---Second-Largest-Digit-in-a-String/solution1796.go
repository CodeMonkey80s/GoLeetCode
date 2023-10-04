package solution1796

import (
	"sort"
)

// ============================================================================
// 1796. Second Largest Digit in a String
// URL: https://leetcode.com/problems/second-largest-digit-in-a-string/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_secondHighest-24    	 9679000	       119.6 ns/op	      96 B/op	       5 allocs/op
	PASS
	ok  	_/home/mp/Development/learning/LeetCode/1796---Second-Largest-Digit-in-a-String	1.285

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
	sort.Sort(sort.Reverse(sort.IntSlice(st)))
	if len(st) > 1 {
		return st[1]
	}
	return -1
}
