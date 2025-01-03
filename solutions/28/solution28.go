package solution28

import "strings"

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/28---Find-the-index-of-the-First-Occurrence-in-a-String
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_strStr
	Benchmark_strStr-24    	305080532	         4.224 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if i+len(needle) > len(haystack) {
			return -1
		}
		if s := haystack[i : i+len(needle)]; s == needle {
			return i
		}
	}
	return -1
}

func strStr_stdlib(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
