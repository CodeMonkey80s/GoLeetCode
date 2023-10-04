package solution14

// ============================================================================
// 14. Longest Common Prefix
// URL: https://leetcode.com/problems/longest-common-prefix/
// ============================================================================

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	prefix := make([]byte, 0, length)
	var char byte
	var c, i, l int
	for {
		c = 0
		for _, str := range strs {
			l = len(str)
			if i >= l {
				return string(prefix)
			}
			if char == 0 {
				char = str[i]
			}
			if char == str[i] {
				c++
			}
		}
		if c != length {
			break
		}
		prefix = append(prefix, char)
		char = 0
		i++
		if i > l {
			break
		}
	}
	return string(prefix)
}
