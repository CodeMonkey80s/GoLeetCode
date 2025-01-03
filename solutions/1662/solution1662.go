package solution1662

import "strings"

// ============================================================================
// 1662. Check If Two String Arrays are Equivalent
// URL: https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent
// ============================================================================

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	s1 := strings.Join(word1, "")
	s2 := strings.Join(word2, "")
	return s1 == s2
}
