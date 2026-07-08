package solution3945

import (
	"strconv"
)

func digitFrequencyScore(n int) int {
	freq := make(map[rune]int)
	for _, d := range strconv.Itoa(n) {
		freq[d-'0']++
	}

	var score int
	for d, f := range freq {
		score += int(d) * f
	}

	return score
}
