package solution2864

// ============================================================================
// 2864. Maximum Odd Binary Number
// URL: https://leetcode.com/problems/maximum-odd-binary-number/
// ============================================================================

/*
	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2864---Maximum-Odd-Binary-Number
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_maximumOddBinaryNumber-24    	96572155	        14.25 ns/op	       3 B/op	       1 allocs/op
	PASS

*/

func maximumOddBinaryNumber(s string) string {
	ans := make([]byte, len(s))
	for i := range ans {
		ans[i] = '0'
	}
	c := 0
	for _, char := range s {
		if char == '1' {
			c++
		}
	}
	ans[len(s)-1] = '1'
	for i := 1; i < c; i++ {
		ans[i-1] = '1'
	}
	return string(ans)
}
