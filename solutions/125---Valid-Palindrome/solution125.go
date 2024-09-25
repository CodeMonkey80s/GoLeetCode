package solution125

import "strings"

// ============================================================================
// 125. Valid Palindrome
// URL: https://leetcode.com/problems/valid-palindrome/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/125---Valid-Palindrome
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_isPalindromeV1
	Benchmark_isPalindromeV1-24    	 8927130	       134.1 ns/op	      56 B/op	       3 allocs/op
	Benchmark_isPalindromeV2
	Benchmark_isPalindromeV2-24    	24183046	        49.66 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func isPalindromeV2(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return true
	}
	isValidChar := func(ch rune) (rune, bool) {
		switch {
		case ch >= 'A' && ch <= 'Z':
			return ch - 65 + 97, true
		case ch >= 'a' && ch <= 'z':
			return ch, true
		case ch >= '0' && ch <= '9':
			return ch, true
		default:
			return 0, false
		}
	}
	a := 0
	b := len(s) - 1
	for {
		if b <= a {
			break
		}
		ch1, ok1 := isValidChar(rune(s[a]))
		ch2, ok2 := isValidChar(rune(s[b]))
		if ok1 && ok2 {
			if ch1 != ch2 {
				return false
			}
			a++
			b--
			continue
		}
		if !ok1 {
			a++
			continue
		}
		b--
	}

	return true
}

func isPalindromeV1(s string) bool {
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
