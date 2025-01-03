package solution3258

// ============================================================================
// 485. Max Consecutive Ones
// URL: https://leetcode.com/problems/max-consecutive-ones/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3258---Count-Substrings-That-Satisfy-K-Constraint-I
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countKConstraintSubstrings
	Benchmark_countKConstraintSubstrings-24    	22556546	        83.19 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func countKConstraintSubstrings(s string, k int) int {
	var ans, v1, v2 int
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			v1 = 0
			v2 = 0
			for n := 0; n < j-i; n++ {
				switch s[i+n] {
				case '1':
					v1++
				case '0':
					v2++
				}
			}
			if v1 <= k || v2 <= k {
				ans++
			}
		}
	}

	return ans
}
