package solution2309

import "fmt"

// ============================================================================
// 2309. Greatest English Letter in Upper and Lower Case
// URL: https://leetcode.com/problems/greatest-english-letter-in-upper-and-lower-case/
// ============================================================================

func greatestLetter(s string) string {
	var val byte
	freqUp := make([]byte, 27)
	freqDo := make([]byte, 27)
	for _, letter := range s {
		val = byte(letter)
		if 'a' <= letter && letter <= 'z' {
			val -= 'a'
			freqDo[val]++
		} else if 'A' <= letter && letter <= 'Z' {
			val -= 'A'
			freqUp[val]++
		}
	}
	fmt.Printf("freqDo: %v\n", freqDo)
	fmt.Printf("freqUp: %v\n", freqUp)
	for i := len(freqDo) - 1; i >= 0; i-- {
		if freqDo[i] >= 1 && freqUp[i] >= 1 {
			return string(byte(i) + 'A')
		}
	}
	return ""
}
