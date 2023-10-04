package solution434

import "strings"

// ============================================================================
// 434. Number of Segments in a String
// URL: https://leetcode.com/problems/number-of-segments-in-a-string/
// ============================================================================

func countSegments(s string) int {
	return len(strings.Fields(s))
}
