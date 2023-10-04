package solution482

import (
	"strings"
)

// ============================================================================
// 482. License Key Formatting
// URL: https://leetcode.com/problems/license-key-formatting/
// ============================================================================

func licenseKeyFormatting(s string, k int) string {
	sl := make([]byte, 0, len(s))
	upper := strings.ToUpper(s)
	d := 0
	for i := len(upper) - 1; i >= 0; i-- {
		ch := byte(upper[i])
		if ch != '-' {
			sl = append(sl, ch)
			d++
		}
		if d == k && i > 0 {
			sl = append(sl, '-')
			d = 0
		}
	}
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	if len(sl) > 1 && sl[0] == '-' {
		return string(sl[1:])
	}
	return string(sl)
}
