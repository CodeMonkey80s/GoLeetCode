package solution451

// ============================================================================
// 451. Sort Characters By Frequency
// URL: https://leetcode.com/problems/sort-characters-by-frequency/
// ============================================================================

import "slices"

func frequencySort(s string) string {

	freq := make(map[rune]int)
	for _, ru := range s {
		freq[ru]++
	}

	runes := []rune(s)

	slices.SortFunc(runes, func(a rune, b rune) int {
		switch {
		case freq[a] == freq[b]:
			return int(a) - int(b)
		case freq[a] > freq[b]:
			return -1
		default:
			return 0
		}
	})

	return string(runes)
}
