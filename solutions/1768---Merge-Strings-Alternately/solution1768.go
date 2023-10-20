package solution1768

import "strings"

// ============================================================================
// 1768. Merge Strings Alternately
// URL: https://leetcode.com/problems/merge-strings-alternately/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1768---Merge-Strings-Alternately
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_mergeAlternately-24    	65478508	        18.57 ns/op	       8 B/op	       1 allocs/op
	PASS

*/

func mergeAlternately(word1 string, word2 string) string {
	sb := strings.Builder{}
	size1 := len(word1)
	size2 := len(word2)
	size := size1
	if size2 > size1 {
		size = size2
	}
	for i := 0; i < size; i++ {
		if i < size1 {
			sb.WriteByte(word1[i])
		}
		if i < size2 {
			sb.WriteByte(word2[i])
		}
	}
	return sb.String()
}
