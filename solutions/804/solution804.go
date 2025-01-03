package solution804

// ============================================================================
// 804. Unique Morse Code Words
// URL: https://leetcode.com/problems/unique-morse-code-words/
// ============================================================================

func uniqueMorseRepresentations(words []string) int {
	sl := []string{
		".-",
		"-...",
		"-.-.",
		"-..",
		".",
		"..-.",
		"--.",
		"....",
		"..",
		".---",
		"-.-",
		".-..",
		"--",
		"-.",
		"---",
		".--.",
		"--.-",
		".-.",
		"...",
		"-",
		"..-",
		"...-",
		".--",
		"-..-",
		"-.--",
		"--..",
	}
	m := make(map[string]int)
	output := ""
	for _, word := range words {
		output = ""
		for _, ch := range word {
			ch := byte(ch) - 97
			output += sl[ch]
		}
		_,
			ok := m[output]
		if ok {
			m[output]++
		} else {
			m[output] = 1
		}
	}
	return len(m)
}
