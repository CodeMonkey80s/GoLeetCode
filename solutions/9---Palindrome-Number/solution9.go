package solution9

import (
	"strconv"
)

// ============================================================================
// 9. Palindrome Number
// URL: https://leetcode.com/problems/palindrome-number/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/9---Palindrome-Number
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_isPalindrome
	Benchmark_isPalindrome-24    	48831478	        20.94 ns/op	       3 B/op	       1 allocs/op
	PASS
*/

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
