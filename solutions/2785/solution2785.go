package solution2785

// ============================================================================
// 2785. Sort Vowels in a String
// URL: https://leetcode.com/problems/sort-vowels-in-a-string/
// ============================================================================

import (
	"slices"
	"strings"
)

func sortVowels(s string) string {

	const (
		chars = "AEIOUaeiou"
	)

	var vowels []rune
	for _, ru := range s {
		if strings.ContainsRune(chars, ru) {
			vowels = append(vowels, ru)
		}
	}

	slices.Sort(vowels)

	var j int
	var sb strings.Builder
	for i, ru := range s {
		if !strings.ContainsRune(chars, ru) {
			sb.WriteRune(rune(s[i]))
		} else {
			sb.WriteRune(vowels[j])
			j++
		}
	}

	return sb.String()
}
