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
	Benchmark_greatestLetter-24    	46853886	        25.69 ns/op	       4 B/op	       1 allocs/op
	PASS

*/

func greatestLetter(s string) string {
	var val byte
	freqUp := make([]byte, 27)
	freqDo := make([]byte, 27)
	for _, letter := range s {
		val = byte(letter)
		if 'a' <= letter && letter <= 'z' {
			val -= 'a'
			freqDo[val]++
		} else if 'A' <= letter && letter <= 'Z' {
			val -= 'A'
			freqUp[val]++
		}
	}
	for i := len(freqDo) - 1; i >= 0; i-- {
		if freqDo[i] >= 1 && freqUp[i] >= 1 {
			return string(byte(i) + 'A')
		}
	}
	return ""
}
