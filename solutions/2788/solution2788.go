package solution2788

import "strings"

// ============================================================================
// 2788. Split Strings by Separator
// URL: https://leetcode.com/problems/split-strings-by-separator/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2788---Split-Strings-by-Separator
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_splitWords-24    	 5995474	       218.5 ns/op	     336 B/op	       7 allocs/op
	PASS

*/

func splitWordsBySeparator(words []string, separator byte) []string {
	ans := make([]string, 0)
	for _, str := range words {
		sl := strings.Split(str, string(separator))
		for _, w := range sl {
			if len(w) > 0 {
				ans = append(ans, w)
			}
		}
	}
	return ans
}
