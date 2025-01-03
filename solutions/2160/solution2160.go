package solution2160

import (
	"cmp"
	"slices"
	"strconv"
)

// ============================================================================
// 2160. Minimum Sum of Four Digit Number After Splitting Digits
// URL: https://leetcode.com/problems/minimum-sum-of-four-digit-number-after-splitting-digits/
// ============================================================================

/*

	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2160---Minimum-Sum-of-Four-Digit-Number-After-Splitting-Digits
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_minimumSum
	Benchmark_minimumSum-24    	24084129	        45.62 ns/op	       4 B/op	       1 allocs/op
	PASS

*/

func minimumSum(num int) int {
	digits := strconv.Itoa(num)

	sl := []byte(digits)

	slices.SortFunc(sl, func(a byte, b byte) int {
		return cmp.Compare(a, b)
	})

	num1 := 10*(sl[0]-'0') + sl[2] - '0'
	num2 := 10*(sl[1]-'0') + sl[3] - '0'

	return int(num1 + num2)
}
