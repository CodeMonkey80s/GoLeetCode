package solution9

import (
	"strconv"
)

// ============================================================================
// 9. Palindrome Number
// URL: https://leetcode.com/problems/palindrome-number/
// ============================================================================

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
