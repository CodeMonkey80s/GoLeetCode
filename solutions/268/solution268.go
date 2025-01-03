package solution268

// ============================================================================
// 268. Missing Number
// URL: https://leetcode.com/problems/missing-number/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/268---Missing-Number
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_missingNumber-24    	904036344	         1.326 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func missingNumber(nums []int) int {
	s1, s2 := 0, 0
	size := len(nums)
	for i := 0; i <= size; i++ {
		s1 += i
		if i < size {
			s2 += nums[i]
		}
	}
	return s1 - s2
}
