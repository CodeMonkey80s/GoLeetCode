package solution151

import "strings"

// ============================================================================
// 151. Reverse Words in a String
// URL: https://leetcode.com/problems/reverse-words-in-a-string/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/151---Reverse-Words-in-a-String
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	BenchmarkReverseWordsV2
	BenchmarkReverseWordsV2-24    	11085066	       124.4 ns/op	      88 B/op	       3 allocs/op
	BenchmarkReverseWordsV1
	BenchmarkReverseWordsV1-24    	 5831085	       205.6 ns/op	     128 B/op	       7 allocs/op
	PASS
*/

func reverseWordsV2(s string) string {
	var sb strings.Builder
	words := strings.Fields(s)
	for i := len(words) - 1; i >= 0; i-- {
		sb.WriteString(words[i])
		if i > 0 {
			sb.WriteRune(' ')
		}
	}
	return sb.String()
}

func reverseWordsV1(s string) string {
	words := strings.Fields(s)
	output := ""
	for i := len(words) - 1; i >= 0; i-- {
		output += words[i]
		if i > 0 {
			output += " "
		}
	}
	return output
}
