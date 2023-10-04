package solution383

// ============================================================================
// 383. Ransom Note
// URL: https://leetcode.com/problems/ransom-note/
// ============================================================================

func canConstruct(ransomNote string, magazine string) bool {
	count1 := len(ransomNote)
	m := make(map[byte]int)
	for _, v := range magazine {
		ch := byte(v)
		_, ok := m[ch]
		if ok {
			m[ch]++
		} else {
			m[ch] = 1
		}
	}
	for i := 0; i < count1; i++ {
		ch := byte(ransomNote[i])
		_, ok := m[ch]
		if ok {
			m[ch]--
			if m[ch] == 0 {
				delete(m, ch)
			}
		} else {
			return false
		}
	}
	if len(m) == 0 {
		return true
	}
	return len(m) != 0
}
