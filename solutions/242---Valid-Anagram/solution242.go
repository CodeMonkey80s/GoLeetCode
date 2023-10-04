package solution242

// ============================================================================
// 242. Valid Anagram
// URL: https://leetcode.com/problems/valid-anagram/description/
// ============================================================================

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var a, b rune = 0, 0
	m := make(map[rune]int)
	for _, v := range s {
		a = a ^ v
		m[v]++
	}
	for _, v := range t {
		b = b ^ v
		_, ok := m[v]
		if !ok {
			return false
		}
		m[v]--
		if m[v] < 0 {
			return false
		}
	}
	if a == 1 && b == 1 && len(m) == 0 {
		return true
	}
	if a == 0 && b == 0 {
		return false
	}
	return a == b
}
