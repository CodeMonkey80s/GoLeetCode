package solution2114

import "strings"

// ============================================================================
// 2114. Maximum Number of Words Found in Sentences
// URL: https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2114---Maximum-Number-of-Words-Found-in-Sentences
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_mostWordsFound-24                  	51084126	        23.66 ns/op	       0 B/op	       0 allocs/op
	Benchmark_mostWordsFound_strings_count-24    	100000000	        17.88 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func mostWordsFound(sentences []string) int {
	ans := 0
	words := 0
	for _, sentence := range sentences {
		words = 0
		for _, char := range sentence {
			if char == ' ' {
				words++
			}
		}
		words++
		if words > ans {
			ans = words
		}
	}
	return ans
}

func mostWordsFound_strings_count(sentences []string) int {
	ans := 0
	words := 0
	for _, sentence := range sentences {
		words = strings.Count(sentence, " ") + 1
		if words > ans {
			ans = words
		}
	}
	return ans
}
