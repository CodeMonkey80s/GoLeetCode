package solution2011

// ============================================================================
// 2011. Final Value of Variable After Performing Operations
// URL: https://leetcode.com/problems/final-value-of-variable-after-performing-operations/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2011---Final-Value-of-Variable-After-Performing-Operations
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_finalValue-24    	1000000000	         1.138 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func finalValueAfterOperations(operations []string) int {
	ans := 0
	for _, operation := range operations {
		ch := operation[1]
		switch ch {
		case '+':
			ans++
		case '-':
			ans--
		}
	}
	return ans
}
