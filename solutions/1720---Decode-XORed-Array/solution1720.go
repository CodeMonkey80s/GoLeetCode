package solution1720

// ============================================================================
// 1720. Decode XORed Array
// URL: https://leetcode.com/problems/decode-xored-array/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1720---Decode-XORed-Array
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_decode-24    	71002912	        14.12 ns/op	      32 B/op	       1 allocs/op
	PASS

*/

func decode(encoded []int, first int) []int {
	ans := make([]int, len(encoded)+1)
	ans[0] = first
	for i := range encoded {
		ans[i+1] = encoded[i] ^ ans[i]
	}
	return ans
}
