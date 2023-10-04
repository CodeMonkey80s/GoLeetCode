package solution459

import "strings"

// ============================================================================
// 459. Repeated Subscting Pattern
// URL: https://leetcode.com/problems/repeated-substring-pattern/
// ============================================================================

func repeatedSubstringPattern(s string) bool {
	sub := ""
	for i, v := range s {
		sub += string(v)
		out := strings.ReplaceAll(s, sub, "")
		if out == "" && i < len(s)-1 {
			return true
		}
	}
	return false
}
