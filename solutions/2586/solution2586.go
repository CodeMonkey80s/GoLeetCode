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
	Benchmark_vowelStrings-24            	315741366	         3.793 ns/op	       0 B/op	       0 allocs/op
	Benchmark_vowelStrings_bitmask-24    	745920466	         1.608 ns/op	       0 B/op	       0 allocs/op

*/

func vowelStrings(words []string, left int, right int) int {
	isVowel := func(ch byte) bool {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u':
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

func vowelStrings_bitmask(words []string, left int, right int) int {
	ans := 0
	vowels := 0b_100000100000100010001
	for i := left; i <= right; i++ {
		if vowels&(1<<(words[i][0]-'a')) > 0 && vowels&(1<<(words[i][len(words[i])-1]-'a')) > 0 {
			ans++
		}
	}
	return ans
}
