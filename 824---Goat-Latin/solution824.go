package solution824

import "strings"

// ============================================================================
// 824. Goat Latin
// URL: https://leetcode.com/problems/goat-latin/
// ============================================================================

func toGoatLatin(sentence string) string {
	m := map[byte]int{
		'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1,
		'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1,
	}
	output := ""
	words := strings.Fields(sentence)
	for i, word := range words {
		ch := byte(word[0])
		_, ok := m[ch]
		if ok {
			output += word + "ma"
		} else {
			output += word[1:] + string(ch) + "ma"
		}
		output += strings.Repeat("a", i+1)
		if i < len(words)-1 {
			output += " "
		}
	}
	return output
}
