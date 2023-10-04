package solution819

import (
	"strings"
	"unicode"
)

// ============================================================================
// 819. Most Common Word
// URL: https://leetcode.com/problems/most-common-word/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_mostCommonWord-24    	 1000000	      1084 ns/op	    4336 B/op	       5 allocs/op
	PASS

*/

func mostCommonWord(paragraph string, banned []string) string {
	output := ""
	maxval := 0
	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		return unicode.IsSpace(r) || !unicode.IsLetter(r)
	})
	m := make(map[string]int, len(paragraph))
outer:
	for _, word := range words {
		w := strings.ToLower(word)
		for _, bword := range banned {
			if bword == w {
				continue outer
			}
		}
		_, ok := m[w]
		if ok {
			m[w]++
		} else {
			m[w] = 1
		}
		if maxval < m[w] {
			maxval = m[w]
			output = w
		}
	}
	return output
}
