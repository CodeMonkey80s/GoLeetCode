package solution345

// ============================================================================
// 345. Reverse Vowels of a String
// URL: https://leetcode.com/problems/reverse-vowels-of-a-string/
// ============================================================================

func reverseVowels(s string) string {
	var (
		count = len(s)
		a     = 0
		b     = count - 1
		ch1   byte
		ch2   byte
	)
	buf := []byte(s)
	m := map[byte]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
		'A': 1,
		'E': 1,
		'I': 1,
		'O': 1,
		'U': 1,
	}
	for i := 0; i < count; i++ {
		ch1 = buf[a]
		ch2 = buf[b]
		if a > b {
			return string(buf)
		}
		if (ch1 < 'A' || ch1 > 'Z') && (ch1 < 'a' || ch1 > 'z') {
			a++
			continue
		}
		if (ch2 < 'A' || ch2 > 'Z') && (ch2 < 'a' || ch2 > 'z') {
			b--
			continue
		}
		_, ok1 := m[ch1]
		_, ok2 := m[ch2]
		if ok1 && ok2 {
			buf[a], buf[b] = buf[b], buf[a]
			a++
			b--
			continue
		}
		if !ok1 {
			a++
		}
		if !ok2 {
			b--
		}
	}
	return string(buf)
}
