package solution1002

import (
	"bytes"
	"sort"
)

// ============================================================================
// 1002. Find Common Characters
// URL: https://leetcode.com/problems/find-common-characters/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_commmonChars-24    	 5433933	       259.2 ns/op	     256 B/op	      11 allocs/op
	PASS

*/

func commonChars(words []string) []string {
	out := make([]string, 0)
	s := make([][]byte, 0, len(words))
	for _, word := range words {
		s = append(s, []byte(word))
	}
	for _, ch := range words[0] {
		found := 0
		for j := 0; j < len(s); j++ {
			ok := bytes.IndexByte(s[j], byte(ch))
			if ok != -1 {
				s[j][ok] = 'X'
				found++
			}
		}
		if found > 0 && found == len(words) {
			out = append(out, string(ch))
		}
	}
	sort.Strings(out)
	return out
}
