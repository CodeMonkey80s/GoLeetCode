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
	var a, b byte
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i < length; i++ {
		a = s[i]
		b = s[length-i-1]
		if a != b {
			return false
		}
	}
	return true
}
