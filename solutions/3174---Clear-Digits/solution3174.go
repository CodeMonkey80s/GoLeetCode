package solution3174

// ============================================================================
// 3174. Clear Digits
// URL: https://leetcode.com/problems/clear-digits/description/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3174---Clear-Digits
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_ClearDigits
	Benchmark_ClearDigits-24    	494588084	         2.484 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func clearDigits(s string) string {
loop:
	for {
		found := false
		for i := 0; i < len(s)-1; i++ {
			a := i
			b := i + 1
			if s[a] >= 'a' && s[a] <= 'z' && s[b] >= '0' && s[b] <= '9' {
				s = s[:a] + s[b+1:]
				found = true
			}
		}
		if !found {
			break loop
		}
		if len(s) == 0 {
			break
		}
	}

	return s
}
