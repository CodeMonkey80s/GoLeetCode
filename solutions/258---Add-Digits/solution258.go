package solution258

// ============================================================================
// 258. Add Digits
// URL: https://leetcode.com/problems/add-digits/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/258---Add-Digits
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_addDigits
	Benchmark_addDigits-24    	207767319	         5.789 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func addDigits(num int) int {
	i := 0
	sum := 0
	for {
		for num > 0 {
			sum = sum + num%10
			num = num / 10
			i++
		}
		num = sum
		if num <= 9 {
			break
		}
		sum = 0
	}
	return sum
}
