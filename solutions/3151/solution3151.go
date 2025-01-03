package solution3151

// ============================================================================
// 3151.Special Array I
// URL: https://leetcode.com/problems/special-array-i/description/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3151---Special-Array-I
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_isArraySpecial-24    	699798559	         3.332 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func isArraySpecial(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i]%2 == 0 && nums[i+1]%2 == 0 || nums[i]%2 == 1 && nums[i+1]%2 == 1 {
			return false
		}
	}

	return true
}
