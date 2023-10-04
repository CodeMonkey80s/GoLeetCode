package solution884

import (
	"strings"
)

// ============================================================================
// 884. Uncommon Words from Two Sentences
// URL: https://leetcode.com/problems/uncommon-words-from-two-sentences/
// ============================================================================

func uncommonFromSentences(s1 string, s2 string) []string {
	sl := make([]string, 0)
	m := make(map[string]int)
	words1 := strings.Fields(s1)
	for _, word := range words1 {
		_, ok := m[word]
		if !ok {
			m[word] = 1
		} else {
			m[word]++
		}
	}
	words2 := strings.Fields(s2)
	for _, word := range words2 {
		_, ok := m[word]
		if !ok {
			m[word] = 1
		} else {
			m[word]++
		}
	}
	for k, v := range m {
		if v == 1 {
			sl = append(sl, k)
		}
	}
	return sl
}
