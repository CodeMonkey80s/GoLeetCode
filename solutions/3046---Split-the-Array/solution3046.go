package solution3046

// ============================================================================
// 3046. Split the Array
// URL: https://leetcode.com/problems/split-the-array/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3046---Split-the-Array
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_isPossibleToSplit-24    	123187705	         9.736 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func isPossibleToSplit(nums []int) bool {
	freq := make([]int, 101)
	for _, n := range nums {
		freq[n]++
		if freq[n] >= 3 {
			return false
		}
	}

	return true
}
