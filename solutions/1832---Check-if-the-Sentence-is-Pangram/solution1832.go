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
	Benchmark_checkIfPangram-24                     	55014789	        21.27 ns/op	       0 B/op	       0 allocs/op
	Benchmark_checkIfPangram_bit_manipulation-24    	66616840	        17.85 ns/op	       0 B/op	       0 allocs/op
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

func checkIfPangram_bit_manipulation(sentence string) bool {
	memo := int32(0)
	for _, ch := range sentence {
		// a --> 1, b --> 10, c --> 100...
		memo |= 1 << int(ch-'a')
	}
	// It is "00000011111111111111111111111111" in binary format
	return memo == 0x3FFFFFF
}
