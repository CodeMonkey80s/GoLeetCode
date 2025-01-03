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
	Benchmark_sortSentence-24    	 6317894	       197.8 ns/op	     208 B/op	       5 allocs/op
	PASS

*/

func sortSentence(s string) string {
	words := strings.Fields(s)
	ans := make([]string, 0)
	ans = append(ans, words...)
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][len(ans[i])-1] < ans[j][len(ans[j])-1]
	})
	for i, word := range ans {
		ans[i] = word[:len(word)-1]
	}
	return strings.Join(ans, " ")
}
