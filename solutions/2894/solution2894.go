package solution2894

// ============================================================================
// 2894. Divisible and Non-divisible Sums Difference
// URL: https://leetcode.com/problems/divisible-and-non-divisible-sums-difference/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2894---Divisible-and-Non-divisible-Sums-Difference
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_differenceOfSums-24    	201257142	         5.968 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func differenceOfSums(n int, m int) int {
	num1 := 0
	num2 := 0
	for i := 1; i <= n; i++ {
		if i%m != 0 {
			num1 += i
		} else {
			num2 += i
		}
	}
	return num1 - num2
}
