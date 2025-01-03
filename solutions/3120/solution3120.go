package solution3120

// ============================================================================
// 3120. Count the Number of Special Characters I
// URL: https://leetcode.com/problems/count-the-number-of-special-characters-i/description/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3120---Count-the-Number-of-Special-Characters-I
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_isArraySpecial
	Benchmark_isArraySpecial-24    	58985012	        19.48 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

const (
	maxFreq  = 52
	maxChars = 26
)

func numberOfSpecialCharacters(word string) int {
	result := 0
	freq := make([]bool, maxFreq)
	for _, ch := range word {
		switch {
		case ch >= 'a' && ch <= 'z':
			freq[ch-'a'+maxChars] = true
		case ch >= 'A' && ch <= 'Z':
			freq[ch-'A'] = true
		}
	}

	for i := 0; i < maxChars; i++ {
		if freq[i] && freq[i+maxChars] {
			result++
		}
	}

	return result
}
