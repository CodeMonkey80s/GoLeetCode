package solution890

import (
	"fmt"

	"GoLeetCode/util"
)

func findAndReplacePattern(words []string, pattern string) []string {
	var output []string

	m := make(map[byte]byte)
	for _, ch := range pattern {
		m[byte(ch)] = 0
	}

	util.PrintOrderedMap(m)

	m2 := make(map[byte]byte)
	for _, word := range words {
		for _, ch := range word {
			m2[byte(ch)] = m[pattern[len(m2)]]
		}
		fmt.Printf("word: %s\n", word)
		util.PrintOrderedMap(m2)
		clear(m2)
	}

	return output
}
