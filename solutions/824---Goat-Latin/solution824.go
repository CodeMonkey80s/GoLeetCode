package solution824

import (
	"bytes"
	"strings"
)

// ============================================================================
// 824. Goat Latin
// URL: https://leetcode.com/problems/goat-latin/
// ============================================================================

/*
	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/824---Goat-Latin
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_toGoatLatin-24        	 8994768	       162.5 ns/op	     172 B/op	       7 allocs/op
	Benchmark_toGoatLatin_map-24    	 1713097	       714.0 ns/op	     475 B/op	      16 allocs/op
	PASS

*/

func toGoatLatin(sentence string) string {
	output := make([]byte, 0, len(sentence)*2)
	words := strings.Fields(sentence)
	for i, word := range words {
		ch := word[0]
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
			ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
			output = append(output, word...)
			output = append(output, "ma"...)
		} else {
			output = append(output, word[1:]...)
			output = append(output, ch)
			output = append(output, "ma"...)
		}
		output = append(output, bytes.Repeat([]byte("a"), i+1)...)
		if i < len(words)-1 {
			output = append(output, " "...)
		}
	}
	return string(output)
}

func toGoatLatin_map(sentence string) string {
	m := map[byte]int{
		'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1,
		'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1,
	}
	output := ""
	words := strings.Fields(sentence)
	for i, word := range words {
		ch := byte(word[0])
		_, ok := m[ch]
		if ok {
			output += word + "ma"
		} else {
			output += word[1:] + string(ch) + "ma"
		}
		output += strings.Repeat("a", i+1)
		if i < len(words)-1 {
			output += " "
		}
	}
	return output
}
