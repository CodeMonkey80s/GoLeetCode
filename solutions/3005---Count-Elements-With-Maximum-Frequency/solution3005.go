package solution3005

// ============================================================================
// 3005. Count Elements With Maximum Frequency
// URL: https://leetcode.com/problems/count-elements-with-maximum-frequency/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3005---Count-Elements-With-Maximum-Frequency
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_maxFrequencyElements-24    	17878630	        66.96 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func maxFrequencyElements(nums []int) int {
	maxval := 0
	freq := make([]int, 101)
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
		if freq[nums[i]] > maxval {
			maxval = freq[nums[i]]
		}
	}
	result := 0
	for _, v := range freq {
		if v == maxval {
			result += maxval
		}
	}

	return result
}
