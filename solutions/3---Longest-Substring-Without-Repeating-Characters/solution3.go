package solution3

// ============================================================================
// 3. Longest Substring Without Repeating Characters
// URL: https://leetcode.com/problems/longest-substring-without-repeating-characters/submissions/
// ============================================================================

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	longest := []byte{}
	sl := make([]byte, 0)
	var a byte
	for i := 0; i < length; i++ {
		a = s[i]
		sl = append(sl, a)
		if i > 0 {
			for j := 0; j < len(sl)-1; j++ {
				if sl[j] == a {
					sl = sl[j+1:]
				}
			}
		}
		if len(sl) > len(longest) {
			longest = sl
		}
	}
	return len(longest)
}
