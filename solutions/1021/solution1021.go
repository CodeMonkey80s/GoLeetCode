package solution1021

import "strings"

// ============================================================================
// 1021.Remove Outermost Parentheses
// URL: https://leetcode.com/problems/remove-outermost-parentheses/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1021
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_removeOuterParentheses
	Benchmark_removeOuterParentheses-24    	15852952	        71.72 ns/op	      32 B/op	       3 allocs/op
	PASS
*/

func removeOuterParentheses(s string) string {

	var (
		n      int
		sb     strings.Builder
		output string
	)

	sb.Grow(len(s))

	for _, ch := range s {
		switch {
		case ch == '(':
			sb.WriteRune('(')
			n++
		case ch == ')':
			sb.WriteRune(')')
			n--
		}

		if n == 0 {
			output += sb.String()[1 : sb.Len()-1]
			sb.Reset()
		}
	}

	return output
}
