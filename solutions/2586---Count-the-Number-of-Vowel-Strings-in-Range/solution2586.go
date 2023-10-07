package solution2586

// ============================================================================
// 2586. Count the Number of Vowel Strings in Range
// URL: https://leetcode.com/problems/count-the-number-of-vowel-strings-in-range/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2586---Count-the-Number-of-Vowel-Strings-in-Range
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_vowelStrings-24    	315818386	         3.803 ns/op	       0 B/op	       0 allocs/op

*/

func vowelStrings(words []string, left int, right int) int {
	isVowel := func(ch byte) bool {
		switch {
		case ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u':
			return true
		}
		return false
	}
	ans := 0
	for i := 0; i < len(words); i++ {
		if i >= left && i <= right {
			word := words[i]
			a := word[0]
			b := word[len(word)-1]
			if isVowel(a) && isVowel(b) {
				ans++
			}
		}
	}
	return ans
}
