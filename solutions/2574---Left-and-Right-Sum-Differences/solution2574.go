package solution2574

// ============================================================================
// 2574. Left and Right Sum Differences
// URL: https://leetcode.com/problems/left-and-right-sum-differences/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2574---Left-and-Right-Sum-Differences
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_leftRightDifference-24    	40379600	        28.79 ns/op	      32 B/op	       1 allocs/op
	PASS

*/

func leftRightDifference(nums []int) []int {
	ans := make([]int, len(nums))
	sum := 0
	for i := 0; i <= len(nums)-1; i++ {
		sum = 0
		if i > 0 {
			for j := 0; j < i; j++ {
				sum += nums[j]
			}
		}
		if i < len(nums) {
			for j := i + 1; j <= len(nums)-1; j++ {
				sum -= nums[j]
			}
		}
		ans[i] = sum
		if ans[i] < 0 {
			ans[i] = -ans[i]
		}
	}
	return ans
}
