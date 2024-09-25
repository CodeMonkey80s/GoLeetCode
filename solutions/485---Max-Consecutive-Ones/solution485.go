package solution485

// ============================================================================
// 485. Max Consecutive Ones
// URL: https://leetcode.com/problems/max-consecutive-ones/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/485---Max-Consecutive-Ones
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_findMaxConsecutiveOnes
	Benchmark_findMaxConsecutiveOnes-24    	246963822	         4.860 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func findMaxConsecutiveOnes(nums []int) int {
	var c, n int
	for _, v := range nums {
		switch {
		case v == 0:
			c = 0
		case v == 1:
			c++
			if c > n {
				n = c
			}
		}
	}

	return n
}
