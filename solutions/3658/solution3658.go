package solution3658

// ============================================================================
// 3658.GCD Of Odd and Even Sums
// URL: https://leetcode.com/problems/gcd-of-odd-and-even-sums/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3658
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_gcdOfOddEvenSums
	Benchmark_gcdOfOddEvenSums-24    	311563850	         3.790 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func gcdOfOddEvenSums(n int) int {
	var (
		sumOdd  int
		sumEven int
	)

	for i := 1; i <= n<<1; i++ {
		switch {
		case i%2 == 0:
			sumEven += i
		default:
			sumOdd += i
		}
	}

	a := sumEven
	b := sumOdd
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
