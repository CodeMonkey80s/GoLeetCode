package solution3216

import (
	"slices"
)

// ============================================================================
// 3216. Lexicographically Smallest String After a Swap
// URL: https://leetcode.com/problems/lexicographically-smallest-string-after-a-swap/
// ============================================================================

func getSmallestString(s string) string {

	var numbers []string

	numbers = append(numbers, s)

	getString := func(s string, idx int) int {
		buf := []rune(s)
		found := false
		i := 0
		for i = idx; i < len(s); i++ {
			if found {
				continue
			}
			if i < len(s)-1 {
				v1 := s[i] - '0'
				v2 := s[i+1] - '0'
				if v1 == v2 || v1 < v2 {
					continue
				}
				if v1%2 == 0 && v2%2 == 0 {
					buf[i], buf[i+1] = rune(s[i+1]), rune(s[i])
					found = true
					i++
					continue
				}
				if v1%2 != 0 && v2%2 != 0 {
					buf[i], buf[i+1] = rune(s[i+1]), rune(s[i])
					found = true
					i++
					continue
				}
			}
		}
		numbers = append(numbers, string(buf))
		return i
	}

	idx := 0
	for {
		idx = getString(s, idx)
		if idx >= len(s) {
			break
		}
	}

	slices.Sort(numbers)

	return numbers[0]
}
