package solution2942

// ============================================================================
// 2942. Find Words Containing Character
// URL: https://leetcode.com/problems/find-words-containing-character/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2942---Find-Words-Containing-Character
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_findWordsContaining-24    	52804471	        20.92 ns/op	      16 B/op	       1 allocs/op
	PASS
*/

func findWordsContaining(words []string, x byte) []int {
	ids := make([]int, 0, len(words))
outer:
	for id, word := range words {
		for i := 0; i < len(word); i++ {
			j := len(word) - 1
			if word[i] == x || word[j] == x {
				ids = append(ids, id)
				continue outer
			}
			j--
		}
	}

	return ids
}
