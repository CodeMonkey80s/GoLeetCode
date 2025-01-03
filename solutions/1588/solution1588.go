package solution1588

// ============================================================================
// 1588. Sum of All Odd Length Subarrays
// URL: https://leetcode.com/problems/sum-of-all-odd-length-subarrays/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1588---Sum-of-All-Odd-Length-Subarrays
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_sumOddLengthsSubarrays-24    	39795235	        30.20 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func sumOddLengthSubarrays(arr []int) int {
	sumRange := func(a int, b int) int {
		ans := 0
		for i := 0; i <= len(arr)-b; i++ {
			for j := a; j < b; j++ {
				idx := i + j
				if idx >= len(arr) {
					return ans
				}
				ans += arr[idx]
			}
			if a+b >= len(arr) {
				break
			}
		}
		return ans
	}

	ans := 0
	a := 0
	b := 1
	for {
		ans += sumRange(a, b)
		b += 2
		if b > len(arr) {
			break
		}
	}
	return ans
}
