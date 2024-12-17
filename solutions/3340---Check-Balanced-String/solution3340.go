package solution3340

// ============================================================================
// 3340. Check Balanced String
// URL: https://leetcode.com/problems/check-balanced-string/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3340---Check-Balanced-String
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_isBalanced
	Benchmark_isBalanced-24    	893154255	         1.477 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func isBalanced(num string) bool {
	sum1 := 0
	sum2 := 0
	for i, ch := range num {
		v := int(ch - '0')
		switch {
		case i%2 == 0:
			sum1 += v
		case i%2 != 0:
			sum2 += v
		}
	}

	return sum1 == sum2
}
