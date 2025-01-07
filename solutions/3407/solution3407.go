package solution3407

import (
	"strings"
)

// ============================================================================
// 3407. Substring Matching Pattern
// URL: https://leetcode.com/problems/substring-matching-pattern/
// ============================================================================

func hasMatch(s string, p string) bool {

	idx := strings.IndexByte(p, '*')

	prefix := p[:idx]
	suffix := p[idx+1:]

	var foundL bool
	var foundR bool

	idx = 0

	if len(prefix) == 0 {
		foundL = true
	} else {
		for i := 0; i <= len(s)-len(prefix); i++ {
			if prefix == s[i:i+len(prefix)] {
				idx = i + len(prefix)
				foundL = true
				break
			}
		}
		s = s[idx:]
	}

	idx = 0

	if len(suffix) == 0 {
		foundR = true
	} else {
		for i := len(s); i >= len(suffix); i-- {
			if suffix == s[i-len(suffix):i] && i-len(suffix) >= idx {
				foundR = true
				break
			}
		}
	}

	return foundL && foundR
}
