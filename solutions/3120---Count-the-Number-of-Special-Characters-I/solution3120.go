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
	Benchmark_isArraySpecial-24    	85115559	        29.59 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func numberOfSpecialCharacters(word string) int {
	result := 0
	freq := make([]int, 52)
	for _, ch := range word {
		idx := ch
		switch {
		case ch >= 'a' && ch <= 'z':
			idx = ch - 97 + 26
			if freq[idx] == 0 {
				freq[idx]++
			}
		case ch >= 'A' && ch <= 'Z':
			idx = ch - 65
			if freq[idx] == 0 {
				freq[idx]++
			}
		}
	}

	for i := 0; i < 26; i++ {
		if freq[i] >= 1 && freq[i+26] >= 1 {
			result++
		}
	}

	return result
}
