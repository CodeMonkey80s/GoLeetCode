package solution2656

// ============================================================================
// 2656. Maximum Sum With Exactly K Elements
// URL: https://leetcode.com/problems/maximum-sum-with-exactly-k-elements/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2656---Maximum-Sum-With-Exactly-K-Elements
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_maximizeSum-24    	648060909	         1.697 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func maximizeSum(nums []int, k int) int {
	ans := 0
	maxval := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxval {
			maxval = nums[i]
		}
	}
	for i := 1; i <= k; i++ {
		ans += maxval
		maxval++
	}
	return ans
}
