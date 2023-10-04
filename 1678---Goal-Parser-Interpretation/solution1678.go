package solution1678

import "strings"

// ============================================================================
// 1678. Goal Parser Interpretation
// URL: https://leetcode.com/problems/goal-parser-interpretation/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_interpret-24    	20711304	        68.33 ns/op	      16 B/op	       2 allocs/op
	PASS

*/

func interpret(command string) string {
	ans := strings.ReplaceAll(command, "(al)", "al")
	ans = strings.ReplaceAll(ans, "()", "o")
	return ans
}
