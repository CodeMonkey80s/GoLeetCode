package solution290

import (
	"strings"
)

// ============================================================================
// 290. Word Pattern
// URL: https://leetcode.com/problems/word-pattern/
// ============================================================================

func wordPattern(pattern string, s string) bool {
	var (
		m  = make(map[byte]string)
		ch byte
	)
	words := strings.Fields(s)
	if len(pattern) != len(words) {
		return false
	}
	for i, rune := range pattern {
		ch = byte(rune)
		v, ok := m[ch]
		if !ok {
			for _, word := range m {
				if word == words[i] {
					return false
				}
			}
			m[ch] = words[i]
			continue
		}
		if v == words[i] {
			continue
		}
		if v != words[i] {
			return false
		}
	}
	return true
}
