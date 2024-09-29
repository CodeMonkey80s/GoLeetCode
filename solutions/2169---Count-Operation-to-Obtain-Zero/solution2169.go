package solution2169

// ============================================================================
// 2169. Count Operations to Obtain Zero
// URL: https://leetcode.com/problems/count-operations-to-obtain-zero/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2169---Count-Operation-to-Obtain-Zero
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_CountOperations
	Benchmark_CountOperations-24    	250241994	         4.773 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func countOperations(num1 int, num2 int) int {
	var n int
	for num1 > 0 && num2 > 0 {
		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}

		n++
	}

	return n
}
