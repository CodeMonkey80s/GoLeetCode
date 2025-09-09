package solution3471

// ============================================================================
// 3471. Find the Largest Almost Missing Integer
// URL: https://leetcode.com/problems/find-the-largest-almost-missing-integer/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3471
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_maxFreqSum
	Benchmark_maxFreqSum-24    	51815949	        19.65 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func largestInteger(nums []int, k int) int {

	var freq [51]int
	var m int

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		freq[n] += min(i+1, len(nums)-i, k)
	}

	if k == len(nums) {
		return m
	}

	for i := 50; i >= 0; i-- {
		if freq[i] == 1 {
			return i
		}
	}

	return -1
}
