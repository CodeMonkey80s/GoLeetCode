package solution1221

// ============================================================================
// 1221. Split a String in Balanced Strings
// URL: https://leetcode.com/problems/split-a-string-in-balanced-strings/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1221---Split-a-String-in-Balanced-Strings
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_cellsInRange-24    	331438161	         3.608 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func balancedStringSplit(s string) int {
	ans := 0
	freqL := 0
	freqR := 0
	for _, char := range s {
		ch := byte(char)
		switch ch {
		case 'L':
			freqL++
		case 'R':
			freqR++
		}
		if freqL == freqR {
			ans++
			freqL = 0
			freqR = 0
		}
	}
	return ans
}
