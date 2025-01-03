package solution1480

// ============================================================================
// 1480. Running Sum of 1d Array
// URL: https://leetcode.com/problems/running-sum-of-1d-array/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1480---Running-Sum-of-1d-Array
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_runningSum-24    	89136436	        15.59 ns/op	      32 B/op	       1 allocs/op
	PASS

*/

func runningSum(nums []int) []int {
	sum := 0
	ans := make([]int, len(nums))
	for i, v := range nums {
		ans[i] = sum + v
		sum = ans[i]
	}
	return ans
}
