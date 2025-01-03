package solution66

// ============================================================================
// 66. Plus One
// URL: https://leetcode.com/problems/plus-one/
// ============================================================================

func plusOne(digits []int) []int {
	m := len(digits)
	s := make([]int, 0, m+1)
	n := -1
	a := false
	for i := m - 1; i >= 0; i-- {
		if i == m-1 {
			digits[i]++
			a = false
		}
		if n != -1 {
			digits[i]++
			n = -1
			a = false
		}
		if digits[i] == 10 {
			digits[i] = 0
			n = i - 1
			a = true
		}
		s = append(s, digits[i])
		if i == 0 && a {
			s = append(s, 1)
		}
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
