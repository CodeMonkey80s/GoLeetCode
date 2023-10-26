package solution1967

import "strings"

// ============================================================================
// 1967. Number of Strings That Appear as Substrings in Word
// URL: https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1967---Number-of-Strings-That-Appear-as-Substrings-in-Word
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_numOfStrings-24    	113775985	        19.79 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func numOfStrings(patterns []string, word string) int {
	ans := 0
	for _, sub := range patterns {
		if strings.Contains(word, sub) {
			ans++
		}
	}
	return ans
}
