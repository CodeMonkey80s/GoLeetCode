package solution1929

// ============================================================================
// 1929. Concatenation of Array
// URL: https://leetcode.com/problems/concatenation-of-array/
// ============================================================================

/*
	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1929---Concatenation-of-Array
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_getConcatenation-24    	82943154	        16.31 ns/op	      48 B/op	       1 allocs/op
	PASS

*/

func getConcatenation(nums []int) []int {
	ans := make([]int, len(nums)*2)
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[i]
		ans[len(nums)+i] = nums[i]
	}
	return ans
}
