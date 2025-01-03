package solution1486

// ============================================================================
// 1486. XOR Operation in an Array
// URL: https://leetcode.com/problems/xor-operation-in-an-array/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1486---XOR-Operation-in-an-Array
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_xorOperation-24    	1000000000	         1.173 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func xorOperation(n int, start int) int {
	ans := 0
	for i := 0; i < n; i++ {
		ans ^= start + 2*i
	}
	return ans
}
