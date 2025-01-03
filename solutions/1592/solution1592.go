package solution1592

import (
	"strings"
)

// ============================================================================
// 1592. Rearrange Spaces Between Words
// URL: https://leetcode.com/problems/rearrange-spaces-between-words/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_reorderSpaces-24    	 6889672	       193.9 ns/op	     136 B/op	       7 allocs/op
	PASS

*/

func reorderSpaces(text string) string {
	if len(text) <= 1 {
		return text
	}
	sb := strings.Builder{}
	spaces := strings.Count(text, " ")
	words := strings.Fields(text)
	s1 := spaces
	if len(words) > 1 {
		s1 = spaces / (len(words) - 1)
	}
	a := 0
	for i, word := range words {
		sb.WriteString(word)
		if i < len(words)-1 {
			sb.WriteString(strings.Repeat(" ", s1))
			a += s1
		}
	}
	sb.WriteString(strings.Repeat(" ", spaces-a))
	return sb.String()
}
