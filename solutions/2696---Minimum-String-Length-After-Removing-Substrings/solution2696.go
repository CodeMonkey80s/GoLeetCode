package solution2696

import (
	"strings"
)

// ============================================================================
// 2696. Minimum String Length After Removing Substrings
// URL: https://leetcode.com/problems/minimum-string-length-after-removing-substrings/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2696---Minimum-String-Length-After-Removing-Substrings
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_minLength-24           		88411903	        13.69 ns/op	       0 B/op	       0 allocs/op
	Benchmark_minLength_string-24    		35640510	        36.39 ns/op	       4 B/op	       1 allocs/op
	Benchmark_minLength_strings_replace-24  13744791	       100.6 ns/op	      16 B/op	       3 allocs/op
	PASS

*/

func minLength(s string) int {
	buf := []byte(s)
outer:
	for i := 0; i < len(buf)-1; i++ {
		a := buf[i]
		b := buf[i+1]
		switch {
		case a == 'A' && b == 'B':
			buf = append(buf[0:i], buf[i+2:]...)
			goto outer
		case a == 'C' && b == 'D':
			buf = append(buf[0:i], buf[i+2:]...)
			goto outer
		}
	}
	return len(buf)
}

func minLength_string(s string) int {
outer:
	for i := 0; i < len(s)-1; i++ {
		a := s[i]
		b := s[i+1]
		switch {
		case a == 'A' && b == 'B':
			s = s[0:i] + s[i+2:]
			goto outer
		case a == 'C' && b == 'D':
			s = s[0:i] + s[i+2:]
			goto outer
		}
	}
	return len(s)
}

func minLength_strings_replace(s string) int {
	for i := 0; i < len(s)-1; i++ {
		s = strings.ReplaceAll(s, "AB", "")
		s = strings.ReplaceAll(s, "CD", "")
	}
	return len(s)
}
