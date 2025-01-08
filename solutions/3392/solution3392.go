package solution3392

// ============================================================================
// 3392. Count Subarrays of Length Three With a Condition
// URL: https://leetcode.com/problems/count-subarrays-of-length-three-with-a-condition/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3392
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countSubarrays
	Benchmark_countSubarrays-24    	395874664	         3.039 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func countSubarrays(nums []int) int {
	var c int
	for i := 1; i <= len(nums)-2; i++ {
		if (nums[i-1]+nums[i+1])*2 == nums[i] {
			c++
		}
	}

	return c
}
