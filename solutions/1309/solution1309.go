package solution1309

// ============================================================================
// 1309. Decrypt String from Alphabet to Integer Mapping
// URL: https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1309---Decrypt-String-from-Alphabet-to-Integer-Mapping
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_freqAlphabet-24    	19125697	        73.33 ns/op	      12 B/op	       4 allocs/op
	PASS

*/

func freqAlphabets(s string) string {
	ans := ""
	letter := byte(0)
	for i := len(s) - 1; i >= 0; i-- {
		switch {
		case '1' <= s[i] && s[i] <= '9':
			letter = byte(s[i]) - '0' - 1
			ans = string(letter+'a') + ans
		case s[i] == '#':
			d2 := byte(s[i-2]) - '0'
			d1 := byte(s[i-1]) - '0'
			letter = d2*10 + d1
			ans = string(letter+'a'-1) + ans
			i -= 2
		}
	}
	return ans
}
