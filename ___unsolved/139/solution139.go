package solution139

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) bool {

	fmt.Println("-----")
	permutations := permute(wordDict, len(wordDict))
	for i, v := range permutations {
		fmt.Printf("i: %d, %s\n", i, v)
	}
	fmt.Println("-----")

loop:
	for _, word := range wordDict {
		if strings.Contains(s, word) {
			s = strings.Replace(s, word, "", -1)
			goto loop
			// c++
		} else {
			return false
		}
	}

	return true
}

func permute(output []string, n int) []string {
	if n == 0 {
		return output
	}

	for _, v := range output {
		output = permute(append(output, v), n-1)
	}

	return output
}
