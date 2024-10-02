package solution3095

// ============================================================================
// 3095. Shortest Subarray With OR at Least K I
// URL: https://leetcode.com/problems/shortest-subarray-with-or-at-least-k-i/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3095---Shortest-Subarray-With-OR-at-Least-K-I
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_minimumSubarrayLength
	Benchmark_minimumSubarrayLength-24    	86145554	        23.01 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func minimumSubarrayLength(nums []int, k int) int {
	ans := 1<<63 - 1
	hasAnswer := false
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			sum := 0
			for _, v := range nums[i:j] {
				sum |= v
			}
			if sum >= k {
				ans = min(ans, len(nums[i:j]))
				hasAnswer = true
			}
		}
	}

	if hasAnswer {
		return ans
	}

	return -1
}
