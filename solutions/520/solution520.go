package solution520

import (
	"strings"
	"unicode"
)

// ============================================================================
// 520. Detect Capital
// URL: https://leetcode.com/problems/detect-capital/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_detectCapitalUse_std-24              	23715241	        56.01 ns/op	       6 B/op	       2 allocs/op
	Benchmark_detectCapitalUse_first_attempt-24    	487452025	         2.463 ns/op	   0 B/op	       0 allocs/op
	PASS
*/

func detectCapitalUse(word string) bool {
	first := unicode.ToTitle(rune(word[0]))
	title := string(first) + strings.ToLower(word[1:])
	upper := strings.ToUpper(word)
	lower := strings.ToLower(word)
	return word == title || word == upper || word == lower
}

func detectCapitalUse_first_attempt(word string) bool {
	isTitle := false
	isUpper := false
	isLower := false
	for i, v := range word {
		ch := byte(v)
		if ch >= 'A' && ch <= 'Z' {
			if isLower {
				return false
			}
			isUpper = true
			if i == 0 {
				isTitle = true
			} else {
				isTitle = false
			}
		} else {
			isLower = true
			if i >= 1 && isUpper && !isTitle {
				return false
			}
		}
	}
	if isTitle {
		return true
	}
	if isUpper {
		return true
	}
	if isLower {
		return true
	}
	return false
}
