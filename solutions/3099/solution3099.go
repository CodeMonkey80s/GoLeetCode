package solution3099

import "strconv"

// ============================================================================
// 3099. Harshad Number
// URL: https://leetcode.com/problems/harshad-number/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3099---Harshad-Number
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_sumOfTheDigitsOfHarshadNumber-24    	194525509	         6.171 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func sumOfTheDigitsOfHarshadNumber(x int) int {
	sum := 0
	n := strconv.Itoa(x)
	for _, v := range n {
		sum += int(v - '0')
	}

	if x%sum == 0 {
		return sum
	}

	return -1
}
