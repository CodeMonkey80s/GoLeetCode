package solution2309

// ============================================================================
// 2309. Greatest English Letter in Upper and Lower Case
// URL: https://leetcode.com/problems/greatest-english-letter-in-upper-and-lower-case/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2309---Greatest-English-Letter-in-Upper-and-Lower-Case
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_greatestLetter-24    	54209395	        22.88 ns/op	       4 B/op	       1 allocs/op
	PASS

*/

func greatestLetter(s string) string {
	var val byte
	upper := [27]byte{}
	lower := [27]byte{}
	for _, letter := range s {
		val = byte(letter)
		if 'a' <= letter && letter <= 'z' {
			val -= 'a'
			lower[val]++
		} else if 'A' <= letter && letter <= 'Z' {
			val -= 'A'
			upper[val]++
		}
	}
	for i := len(lower) - 1; i >= 0; i-- {
		if lower[i] >= 1 && upper[i] >= 1 {
			return string(byte(i) + 'A')
		}
	}
	return ""
}
