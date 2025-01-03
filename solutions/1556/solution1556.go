package solution1556

import (
	"strconv"
)

// ============================================================================
// 1556. Thousand Separator
// URL: https://leetcode.com/problems/thousand-separator/
// ============================================================================

func thousandSeparator(n int) string {
	s := strconv.Itoa(n)
	if len(s) <= 3 {
		return s
	}
	sl := make([]byte, 0, len(s)+len(s)/4)
	c := 1
	for i := len(s) - 1; i >= 0; i-- {
		if c == 4 {
			sl = append(sl, '.')
			c = 1
		}
		sl = append(sl, s[i])
		c++
	}

	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	return string(sl)
}
