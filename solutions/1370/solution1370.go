package solution1370

import (
	"strings"
)

func sortString(s string) string {

	const (
		maxLetters = 26
	)

	var characters int
	freq := make([]byte, maxLetters)
	for i := 0; i < len(s); i++ {
		ch := s[i] - 'a'
		freq[ch]++
		characters++
	}

	dx := 1
	var idx int
	var sb strings.Builder
	for characters > 0 {

		if freq[idx] > 0 {
			freq[idx]--
			sb.WriteByte(byte(idx + 'a'))
			characters--
		}

		idx += dx
		if idx < 0 {
			idx = 0
			dx = 1
			continue
		}

		if idx > maxLetters-1 {
			idx = maxLetters - 1
			dx = -1
			continue
		}

	}

	return sb.String()
}
