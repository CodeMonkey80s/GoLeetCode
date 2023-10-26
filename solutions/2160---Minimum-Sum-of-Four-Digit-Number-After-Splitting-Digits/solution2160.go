package solution2160

import (
	"sort"
	"strconv"
)

// ============================================================================
// 2160. Minimum Sum of Four Digit Number After Splitting Digits
// URL: https://leetcode.com/problems/minimum-sum-of-four-digit-number-after-splitting-digits/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2160---Minimum-Sum-of-Four-Digit-Number-After-Splitting-Digits
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_minimumSum-24    	10960686	       122.0 ns/op	      72 B/op	       4 allocs/op
	PASS

*/

func minimumSum(num int) int {
	digits := strconv.Itoa(num)
	sl := []byte(digits)
	sort.Slice(
		sl,
		func(i int, j int) bool {
			return sl[i] < sl[j]
		},
	)
	num1 := string(sl[0]) + string(sl[2])
	num2 := string(sl[1]) + string(sl[3])
	v1, err := strconv.Atoi(num1)
	if err != nil {
		panic(err)
	}
	v2, _ := strconv.Atoi(num2)
	if err != nil {
		panic(err)
	}
	return v1 + v2
}
