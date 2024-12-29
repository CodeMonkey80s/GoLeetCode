package solution2221

// ============================================================================
// 2221. Find Triangular Sum of an Array
// URL: https://leetcode.com/problems/find-triangular-sum-of-an-array/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2221
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_triangularSum
	Benchmark_triangularSum-24    	64431428	        18.61 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func triangularSum(nums []int) int {

	for len(nums) > 1 {
		for i := 0; i < len(nums)-1; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}

		nums = nums[:len(nums)-1]
	}

	return nums[0]
}
