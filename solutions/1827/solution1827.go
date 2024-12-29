package solution1827

// ============================================================================
// 1827. Minimum Operations to Make the Array Increasing
// URL: https://leetcode.com/problems/minimum-operations-to-make-the-array-increasing
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1827
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_minOperations
	Benchmark_minOperations-24    	518000917	         4.067 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func minOperations(nums []int) int {

	if len(nums) == 1 {
		return 0
	}

	var diff int
	var operations int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			diff = (nums[i] - nums[i+1]) + 1
			nums[i+1] += diff
			operations += diff
		}
	}

	return operations
}
