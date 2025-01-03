package solution3163

// ============================================================================
// 3163. String Compression III
// URL: https://leetcode.com/problems/string-compression-iii/
// ============================================================================

import (
	"strings"
)

func compressedString(word string) string {

	var i int
	var c int
	var ch rune
	var sb strings.Builder
	ch = rune(word[0])

	sb.Grow(2 * len(word))

loop:
	for {

		switch {
		case i > len(word)-1:
			sb.WriteByte(byte(c + '0'))
			sb.WriteRune(ch)
			break loop
		case c == 9:
			sb.WriteByte(byte(c + '0'))
			sb.WriteRune(ch)
			ch = rune(word[i])
			c = 0
		case rune(word[i]) == ch:
			c++
			i++
		case rune(word[i]) != ch:
			sb.WriteByte(byte(c + '0'))
			sb.WriteRune(ch)
			ch = rune(word[i])
			c = 0
		}
	}

	return sb.String()
}
