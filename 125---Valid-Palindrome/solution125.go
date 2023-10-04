package solution125

import (
	"strings"
)

// ============================================================================
// 125. Valid Palindrome
// URL: https://leetcode.com/problems/valid-palindrome/
// ============================================================================

func isPalindrome(s string) bool {
	var c, ch1, ch2 byte
	var a, b, i int
	length := len(s)
	var sb strings.Builder
	for i = 0; i < length; i++ {
		c = s[i]
		if c >= 'A' && c <= 'Z' {
			c = c - 65 + 97
			sb.WriteByte(c)
			continue
		}
		if c >= 'a' && c <= 'z' {
			sb.WriteByte(c)
			continue
		}
		if c >= '0' && c <= '9' {
			sb.WriteByte(c)
			continue
		}
	}
	s2 := sb.String()
	length = len(s2)
	if length == 1 {
		if s2[0] <= 'a' && s2[0] >= 'z' {
			return false
		}
	}
	for i = 0; i < length; i++ {
		a = i
		b = length - i - 1
		ch1 = s2[a]
		ch2 = s2[b]
		if ch1 != ch2 {
			return false
		}
	}
	return true

}
