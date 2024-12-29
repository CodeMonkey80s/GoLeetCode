package solution1910

// ============================================================================
// 1910. Remove All Occurrences of a Substring
// URL: https://leetcode.com/problems/remove-all-occurrences-of-a-substring/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1910
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_removeOccurrencesV1
	Benchmark_removeOccurrencesV1-24    	 6466885	       180.3 ns/op	      32 B/op	       3 allocs/op
	Benchmark_removeOccurrencesV2
	Benchmark_removeOccurrencesV2-24    	11691258	        94.42 ns/op	      16 B/op	       2 allocs/op
	PASS
*/

import (
	"strings"
)

func removeOccurrencesV2(s string, part string) string {

loop:
	for i := 0; i <= len(s)-len(part); i++ {
		if s[i:i+len(part)] == part {
			s = s[:i] + s[i+len(part):]
			goto loop
		}
	}

	return s
}

func removeOccurrencesV1(s string, part string) string {

	for {
		idx := strings.Index(s, part)
		if idx == -1 {
			break
		}

		s = strings.Replace(s, part, "", 1)
	}

	return s
}
