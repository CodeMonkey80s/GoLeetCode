package solution3019

// ============================================================================
// 3019. Number of Changing Keys
// URL: https://leetcode.com/problems/number-of-changing-keys/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3019---Number-of-Changing-Keys
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countKeyChanges-24    	373075306	         3.215 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func countKeyChanges(s string) int {
	changes := 0
	for i := 0; i < len(s)-1; i++ {
		a := i
		b := i + 1
		ch1 := s[a]
		ch2 := s[b]
		switch {
		case ch1 == ch2:
			continue
		case ch1+32 == ch2+32:
			continue
		case ch1-32 == ch2-32:
			continue
		case ch1-32 == ch2:
			continue
		case ch1 == ch2-32:
			continue
		case ch1+32 == ch2:
			continue
		case ch1 == ch2+32:
			continue
		default:
			changes++
		}
	}

	return changes
}
