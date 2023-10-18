package solution1614

// ============================================================================
// 1614. Maximum Nesting Depth of the Parentheses
// URL: https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1614---Maximum-Nesting-Depth-of-the-Parentheses
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_maxDepth-24    	68548447	        17.46 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func maxDepth(s string) int {
	d := 0
	ans := 0
	for _, v := range s {
		switch {
		case v == '(':
			d++
		case v == ')':
			d--
		}
		if d > ans {
			ans = d
		}
	}
	return ans
}
