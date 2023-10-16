package solution2553

import "strconv"

// ============================================================================
// 2553. Separate the Digits in an Array
// URL: https://leetcode.com/problems/separate-the-digits-in-an-array/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2553---Separate-the-Digits-in-an-Array
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_separateDigits-24    	16290290	        83.92 ns/op	     120 B/op	       4 allocs/op
	PASS

*/

func separateDigits(nums []int) []int {
	ans := make([]int, 0)
	for _, num := range nums {
		digits := strconv.Itoa(num)
		for _, digit := range digits {
			ans = append(ans, int(digit-'0'))
		}
	}
	return ans
}
