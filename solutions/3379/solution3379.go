package solution3379

// ============================================================================
// 3379. Transformed Array
// URL: https://leetcode.com/problems/transformed-array/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3379
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_constructTransformedArray
	Benchmark_constructTransformedArray-24    	53425003	        21.98 ns/op	      32 B/op	       1 allocs/op
	PASS
*/

func constructTransformedArray(nums []int) []int {

	result := make([]int, len(nums))

	var idx int
	for i := range nums {
		idx = (i + nums[i]) % len(nums)
		if idx < 0 {
			idx = len(nums) + idx
		}
		result[i] = nums[idx]
	}

	return result
}
