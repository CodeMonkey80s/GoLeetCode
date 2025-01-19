package solution2442

import (
	"slices"
	"strconv"
	"strings"
)

func countDistinctIntegers(nums []int) int {

	freq := make(map[string]int)

	rev := func(s string) string {
		sb := []byte(s)
		slices.Reverse(sb)
		return strings.TrimLeft(string(sb), "0")
	}

	for _, num := range nums {
		s1 := strconv.Itoa(num)
		s2 := rev(s1)
		freq[s1]++
		freq[s2]++
	}

	return len(freq)
}
