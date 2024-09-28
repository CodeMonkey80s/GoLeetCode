package solution3264

// ============================================================================
// 3264. Final Array State After K Multiplication Operations I
// URL: https://leetcode.com/problems/final-array-state-after-k-multiplication-operations-i/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3264---Final-Array-State-After-K-Multiplication-Operations-I
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	BenchmarkGetFinalState
	BenchmarkGetFinalState-24    	41993293	        28.47 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func getFinalState(nums []int, k int, multiplier int) []int {
	var idx, m int
	for i := 0; i < k; i++ {
		idx = 0
		m = 1<<63 - 1
		for j, v := range nums {
			if v < m {
				m = v
				idx = j
			}
		}
		nums[idx] = nums[idx] * multiplier
	}

	return nums
}
