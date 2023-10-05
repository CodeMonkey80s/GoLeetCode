package solution1624

// ============================================================================
// 1624. Largest Substring Between Two Equal Characters
// URL: https://leetcode.com/problems/largest-substring-between-two-equal-characters/
// ============================================================================

func maxLengthBetweenEqualCharacters(s string) int {
	ans := -1
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= 0; j-- {
			if s[i] == s[j] {
				diff := j - i - 1
				if diff > ans {
					ans = diff
				}
			}
		}
	}
	return ans
}
