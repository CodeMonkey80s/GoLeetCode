package solution916

import (
	"fmt"
	"maps"
	"slices"
)

func wordSubsets(words1 []string, words2 []string) []string {

	freq := make(map[rune]int)
	maxFreq := make(map[rune]int)
	output := make(map[string]struct{})

	for _, w := range words2 {
		for _, ru := range w {
			maxFreq[ru]++
		}

		fmt.Printf("maxFreq: %v\n", maxFreq)

		for _, word := range words1 {
			fmt.Printf("checking: %q\n", word)
			for _, ru := range word {
				freq[ru]++
			}

			fmt.Printf("freq: %v\n", freq)

			c := 0
			for ru, count := range maxFreq {
				if count <= freq[ru] {
					c++
				} else {
					break
				}
			}

			if c == len(maxFreq) {
				output[word] = struct{}{}
				fmt.Printf("adding: %v\n", word)
			}

			clear(freq)
		}

		clear(maxFreq)
	}

	return slices.Collect(maps.Keys(output))
}
