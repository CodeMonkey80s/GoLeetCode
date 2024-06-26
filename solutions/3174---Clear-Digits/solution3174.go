package solution3174

// ============================================================================
// 3174. Clear Digits
// URL: https://leetcode.com/problems/clear-digits/description/
// ============================================================================

func clearDigits(s string) string {
loop:
	for {
		found := false
		for i := 0; i < len(s)-1; i++ {
			a := i
			b := i + 1
			if s[a] >= 'a' && s[a] <= 'z' && s[b] >= '0' && s[b] <= '9' {
				s = s[:a] + s[b+1:]
				found = true
			}
		}
		if !found {
			break loop
		}
		if len(s) == 0 {
			break
		}
	}

	return s
}
