package solution205

// ============================================================================
// 205. Isomorphic Strings
// URL: https://leetcode.com/problems/isomorphic-strings/
// ============================================================================

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]byte)
	var ch1, ch2, v1, v2, ch byte
	var ok1, ok2 bool
	for i := 0; i < len(s); i++ {
		ch1 = byte(s[i])
		ch2 = byte(t[i])
		v1, ok1 = m[ch1]
		if ok1 {
			if v1 != ch2 {
				return false
			}
		} else {
			v2, ok2 = m[ch2]
			if ok2 {
				if v2 == ch2 {
					return false
				}
			}
			for _, ch = range m {
				if ch == ch2 {
					return false
				}
			}
		}
		m[ch1] = ch2
	}
	return true
}
