package solution2785

// ============================================================================
// 2785. Sort Vowels in a String
// URL: https://leetcode.com/problems/sort-vowels-in-a-string/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2785
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_sortVowelsV2
	Benchmark_sortVowelsV2-24    	 9544694	       113.7 ns/op	      16 B/op	       2 allocs/op
	Benchmark_sortVowelsV1
	Benchmark_sortVowelsV1-24    	 7979403	       129.8 ns/op	      32 B/op	       3 allocs/op
	PASS
*/

import (
	"slices"
	"strings"
)

func sortVowelsV2(s string) string {

	const (
		chars = "AEIOUaeiou"
	)

	var vowels []byte
	for _, ru := range s {
		if strings.ContainsRune(chars, ru) {
			vowels = append(vowels, byte(ru))
		}
	}

	slices.Sort(vowels)

	b := []byte(s)

	var j int
	for i, ru := range s {
		if strings.ContainsRune(chars, ru) {
			b[i] = vowels[j]
			j++
		}
	}

	return string(b)
}

func sortVowelsV1(s string) string {

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
