package solution3550

// ============================================================================
// 3550. Smallest Index With Digit Sum Equal to Index
// URL: https://leetcode.com/problems/smallest-index-with-digit-sum-equal-to-index/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3550
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_smallestInex
	Benchmark_smallestInex-24    	418159394	         2.820 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func smallestIndex(nums []int) int {
	for i, num := range nums {
		var sum int
		for num > 0 {
			sum += num % 10
			num = num / 10
		}

		if sum == i {
			return i
		}
	}

	return -1
}
