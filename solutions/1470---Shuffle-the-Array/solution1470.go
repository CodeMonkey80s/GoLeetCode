package solution1470

// ============================================================================
// 1470. Shuffle Array
// URL: https://leetcode.com/problems/shuffle-the-array/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1470---Shuffle-the-Array
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_shuffle-24    	87454441	        17.18 ns/op	      48 B/op	       1 allocs/op
	PASS

*/

func shuffle(nums []int, n int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < n; i++ {
		ans[2*i] = nums[i]
		ans[(2*i)+1] = nums[n+i]
	}
	return ans
}
