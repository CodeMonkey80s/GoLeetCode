package solution557

import "strings"

// ============================================================================
// 557. Reverse Words in a String III
// URL: https://leetcode.com/problems/reverse-words-in-a-string-iii
// ============================================================================

/*

	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_reverseWords-24    	 6417997	       190.6 ns/op	     192 B/op	       7 allocs/op
	PASS

*/

func reverseWords(s string) string {
	words := strings.Fields(s)
	sl := make([]string, 0, len(words))
	for _, word := range words {
		r := []byte(word)
		for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		sl = append(sl, string(r))
	}
	return strings.Join(sl, " ")
}
