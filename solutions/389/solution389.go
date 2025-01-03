package solution389

// ============================================================================
// 389. Find the Difference
// URL: https://leetcode.com/problems/find-the-difference/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_findTheDifference-24    	140786538	         8.650 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func findTheDifference(s string, t string) byte {
	if len(s) == 0 {
		return t[0]
	}
	n := 0
	sl := make([]int, 26)
	for i := 0; i < len(t); i++ {
		n = int(t[i]) - 97
		if sl[n] == 0 {
			sl[n] = 1
		} else {
			sl[n] = 0
		}
	}
	for i := 0; i < len(s); i++ {
		n = int(s[i]) - 97
		if sl[n] == 0 {
			sl[n] = 1
		} else {
			sl[n] = 0
		}
	}
	for i := 0; i < len(t); i++ {
		if sl[i] == 1 {
			return byte(i + 97)
		}
	}
	return 'a'
}
