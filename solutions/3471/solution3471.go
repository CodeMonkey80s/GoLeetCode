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

	const (
		maxNumbers = 50
	)

	freq := make([]int, maxNumbers)

	for i := 0; i <= len(nums)-k; i++ {
		for j := 0; j < k; j++ {
			n := nums[i+j]
			freq[n]++
		}
	}

	for i := maxNumbers - 1; i >= 0; i-- {
		switch {
		case freq[i] == 0:
			continue
		case freq[i] == 1:
			return i
		}
	}

	return -1
}
