package solution3310

// ============================================================================
// 3110. Score of String
// URL: https://leetcode.com/problems/score-of-a-string/description/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3110---Score-of-String
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_isArraySpecial-24    	354667786	         3.378 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func scoreOfString(s string) int {
	score := 0
	for i := 0; i < len(s)-1; i++ {
		a := i
		b := i + 1
		v := int(s[a]) - int(s[b])
		if v < 0 {
			v = -v
		}
		score += v
	}

	return score
}
