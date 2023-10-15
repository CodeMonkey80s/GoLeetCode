package solution1313

// ============================================================================
// 1313. Decompress Run-Length Encoded List
// URL: https://leetcode.com/problems/decompress-run-length-encoded-list/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1313---Decompress-Run-Length-Encoded-List
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_decompress-24    	29512189	        45.21 ns/op	      56 B/op	       3 allocs/op
	PASS

*/

func decompressRLElist(nums []int) []int {
	ans := make([]int, 0)
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			ans = append(ans, nums[i+1])
		}
	}
	return ans
}
