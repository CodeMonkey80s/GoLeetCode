package solution1455

import "strings"

// ============================================================================
// 1455. Check If a Word Occurs As a Prefix of Any Word in a Sentence
// URL: https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_isPrefixOfWord-24    	21756010	        65.10 ns/op	      64 B/op	       1 allocs/op
	PASS

*/

func isPrefixOfWord(sentence string, searchWord string) int {
	words := strings.Fields(sentence)
	for i, word := range words {
		if ok := strings.Index(word, searchWord); ok == 0 {
			return i + 1
		}
	}
	return -1
}
