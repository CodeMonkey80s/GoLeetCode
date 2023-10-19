package solution434

import "strings"

// ============================================================================
// 434. Number of Segments in a String
// URL: https://leetcode.com/problems/number-of-segments-in-a-string/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/434---Number-of-Segments-in-a-String
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_missingNumber-24                   	56915252	        20.70 ns/op	       0 B/op	       0 allocs/op
	Benchmark_missingNumber_strings_fields-24    	22404073	        57.33 ns/op	      80 B/op	       1 allocs/op
	PASS

*/

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans := 0
	flag := false
	for _, ch := range s {
		if !flag && ch != ' ' {
			flag = true
			ans++
			continue
		}
		if flag && ch == ' ' {
			flag = false
		}
	}
	return ans
}

func countSegments_strings_fields(s string) int {
	return len(strings.Fields(s))
}
