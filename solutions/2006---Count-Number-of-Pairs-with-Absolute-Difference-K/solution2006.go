package solution2006

import "math"

// ============================================================================
// 2006. Count Number of Pairs With Absolute Difference K
// URL: https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2006---Count-Number-of-Pairs-with-Absolute-Difference-K
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_count-24    	487892126	         2.387 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func countKDifference(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if int(math.Abs(float64(nums[i]-nums[j]))) == k {
				ans++
			}
		}
	}
	return ans
}
