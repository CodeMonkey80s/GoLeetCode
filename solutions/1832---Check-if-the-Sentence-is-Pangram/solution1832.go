package solution1832

// ============================================================================
// 1832. Check if the Sentence Is Pangram
// URL: https://leetcode.com/problems/check-if-the-sentence-is-pangram/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1832---Check-if-the-Sentence-is-Pangram
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_checkIfPangram-24    	56575450	        21.12 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func checkIfPangram(sentence string) bool {
	freq := [26]byte{}
	n := 0
	for _, char := range sentence {
		ch := byte(char) - 'a'
		if freq[ch] == 0 {
			freq[ch] = 1
			n++
		} else {
			freq[ch]++
		}
	}
	return n == 26
}
