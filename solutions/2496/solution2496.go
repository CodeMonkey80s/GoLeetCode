package solution2496

import (
	"strconv"
)

// ============================================================================
// 2496. Maximum Value of a String in an Array
// URL: https://leetcode.com/problems/maximum-value-of-a-string-in-an-array/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2496---Maximum-Value-of-a-String-in-an-Array
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_maximumValue-24    	16113102	        80.53 ns/op	     104 B/op	       4 allocs/op
	PASS

*/

func maximumValue(strs []string) int {
	ans := 0
	for _, word := range strs {
		v, err := strconv.Atoi(word)
		if err != nil {
			if len(word) > ans {
				ans = len(word)
			}
		} else {
			if v > ans {
				ans = v
			}
		}
	}
	return ans
}
