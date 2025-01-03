package solution2778

// ============================================================================
// 2778. Sum of Squares of Special Elements
// URL: https://leetcode.com/problems/sum-of-squares-of-special-elements/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2778---Sum-of-Squares-of-Special-Elements
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_sumOfSquares-24    	1000000000	         0.7297 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func sumOfSquares(nums []int) int {
	ans := 0
	for i, num := range nums {
		if len(nums)%(i+1) == 0 {
			ans += num * num
		}
	}
	return ans
}
