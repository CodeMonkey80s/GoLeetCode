package solution3541

// ============================================================================
// 3541. Find Most Frequent Vowel and Consonant
// URL: https://leetcode.com/problems/find-most-frequent-vowel-and-consonant/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3541
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_maxFreqSum
	Benchmark_maxFreqSum-24    	122230056	         9.824 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func maxFreqSum(s string) int {

	const (
		maxLetters = 26
	)

	freq := make([]byte, maxLetters)

	var maxVowel int
	var maxConsonant int
	for _, ch := range s {
		idx := ch - 'a'
		freq[idx]++
		switch {
		case ch == 'a':
			fallthrough
		case ch == 'e':
			fallthrough
		case ch == 'i':
			fallthrough
		case ch == 'o':
			fallthrough
		case ch == 'u':
			maxVowel = max(int(freq[idx]), maxVowel)
		default:
			maxConsonant = max(int(freq[idx]), maxConsonant)
		}
	}

	return maxVowel + maxConsonant
}
