package solution58

import (
	"strings"
)

// ============================================================================
// 58. Length Of Last Word
// URL: https://leetcode.com/problems/length-of-last-word/
// ============================================================================

func lengthOfLastWord(s string) int {
	length := len(s)
	if length < 1 || length > 10_000 {
		return 0
	}
	c := 0
	word := false
	for i := length - 1; i >= 0; i-- {
		if s[i] == 32 && !word {
			word = false
			c = 0
			continue
		}
		if s[i] == 32 && word {
			return c
		}
		word = true
		c++
	}
	return c
}

func lengthOfLastWord_stringsField(s string) int {
	length := len(s)
	if length < 1 || length > 10_000 {
		return 0
	}
	sl := strings.Fields(s)
	word := sl[len(sl)-1]
	return len(word)
}
