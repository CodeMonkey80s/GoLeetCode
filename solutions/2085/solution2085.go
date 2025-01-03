package solution2085

// ============================================================================
// 2085. Count Common Words With One Occurrence
// URL: https://leetcode.com/problems/count-common-words-with-one-occurrence/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2085---Count-Common-Words-With-One-Occurrence
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countWords-24    	 9253120	       226.6 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func countWords(words1 []string, words2 []string) int {
	ans := 0
	freq1 := make(map[string]int, len(words1))
	freq2 := make(map[string]int, len(words2))
	for _, word := range words1 {
		_, ok := freq1[word]
		if !ok {
			freq1[word] = 1
		} else {
			freq1[word]++
		}
	}
	for _, word := range words2 {
		_, ok := freq2[word]
		if !ok {
			freq2[word] = 1
		} else {
			freq2[word]++
		}
	}
	for k, v1 := range freq1 {
		v2, ok := freq2[k]
		if ok && v1 == 1 && v2 == 1 {
			ans++
		}
	}
	return ans
}
