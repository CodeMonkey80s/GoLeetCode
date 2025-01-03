package solution3136

// ============================================================================
// 3136. Valid Word
// URL: https://leetcode.com/problems/valid-word/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3136---Valid-Word
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	BenchmarkIsValid
	BenchmarkIsValid-24    	133429677	        10.96 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	//          abcdefghijklmnopqrstuvwxyz
	letters := "10001000100000100000100000"

	hasVowel := false
	hasConsonant := false
	for _, ch := range word {
		switch {
		case ch >= 'a' && ch <= 'z':
			if letters[ch-'a'] == '1' {
				hasVowel = true
			} else {
				hasConsonant = true
			}
		case ch >= 'A' && ch <= 'Z':
			if letters[ch-'A'] == '1' {
				hasVowel = true
			} else {
				hasConsonant = true
			}
		case ch >= '0' && ch <= '9':
			continue
		default:
			return false
		}
	}

	return hasVowel && hasConsonant
}
