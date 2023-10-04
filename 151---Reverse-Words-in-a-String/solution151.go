package solution151

import "strings"

// ============================================================================
// 151. Reverse Words in a String
// URL: https://leetcode.com/problems/reverse-words-in-a-string/
// ============================================================================

func reverseWords(s string) string {
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
