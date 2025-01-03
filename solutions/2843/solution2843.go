package solution2843

// ============================================================================
// 2843. Count Symmetric Integers
// URL: https://leetcode.com/problems/count-symmetric-integers/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2843---Count-Symmetric-Integers
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countSymmetricIntegers-24    	 4957764	       223.8 ns/op	       3 B/op	       1 allocs/op
	PASS
*/

import "strconv"

func countSymmetricIntegers(low int, high int) int {
	ans := 0
	for i := low; i <= high; i++ {
		s := strconv.Itoa(i)
		switch {
		case len(s) == 2:
			if s[0] == s[1] {
				ans++
			}
		case len(s) == 4:
			v1 := byte(s[0]) - '0' + s[1] - '0'
			v2 := byte(s[2]) - '0' + s[3] - '0'
			if v1 == v2 {
				ans++
			}
		}
	}
	return ans
}
