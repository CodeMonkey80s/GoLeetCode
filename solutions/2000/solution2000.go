package solution2000

// ============================================================================
// 2000. Reverse Prefix of Word
// URL: https://leetcode.com/problems/reverse-prefix-of-word/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2000---Reverse-Prefix-of-Word
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_reversePrefix-24    	84749766	        16.08 ns/op	       8 B/op	       1 allocs/op
	PASS

*/

import (
	"strings"
)

func reversePrefix(word string, ch byte) string {
	idx := strings.IndexByte(word, ch)
	if idx != -1 {
		rev := func(s []byte, n int) []byte {
			for i, j := 0, n; i < j; i, j = i+1, j-1 {
				s[i], s[j] = s[j], s[i]
			}
			return s
		}
		return string(rev([]byte(word), idx))
	}
	return word
}
