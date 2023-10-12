package solution1859

import (
	"sort"
	"strings"
)

// ============================================================================
// 1859. Sorting the Sentence
// URL: https://leetcode.com/problems/sorting-the-sentence/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1859---Sorting-the-Sentence
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_sortSentence-24    	 4906712	       269.0 ns/op	     256 B/op	       7 allocs/op
	PASS

*/

func sortSentence(s string) string {
	words := strings.Fields(s)
	ans := make([]string, 0)
	for _, word := range words {
		ans = append(ans, word)
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][len(ans[i])-1] < ans[j][len(ans[j])-1]
	})
	for i, word := range ans {
		ans[i] = word[:len(word)-1]
	}
	return strings.Join(ans, " ")
}
