package solution821

// ============================================================================
// 821. Shortest Distance to a Character
// URL: https://leetcode.com/problems/shortest-distance-to-a-character/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_shortestToChar-24    	33337993	        37.71 ns/op	      96 B/op	       1 allocs/op
	PASS

*/

func shortestToChar(s string, c byte) []int {
	a := 0
	b := 0
	sl := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if byte(s[i]) != c {
			a = 0
			b = 0
			for {
				if i-a >= 0 && byte(s[i-a]) == c {
					sl[i] = a
					break
				}
				if i+b < len(s) && byte(s[i+b]) == c {
					sl[i] = b
					break
				}
				a++
				b++
				if a > len(s) && b > len(s) {
					break
				}
			}
		}
	}
	return sl
}
