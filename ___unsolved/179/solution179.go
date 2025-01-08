package solution179

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {

	var numbers []string
	for _, n := range nums {
		num := strconv.Itoa(n)
		numbers = append(numbers, num)
	}

	slices.SortFunc(numbers, func(a, b string) int {
		minlen := min(len(a), len(b))
		var i int
		for i = 0; i < minlen && a[i] == b[i]; i++ {
		}
		switch {
		case i < minlen:
			return int(b[i] - a[i])
		default:
			return len(a[:i]+b) - len(b[:i]+a)
		}
	})

	// n: 09: 001001
	// n: 05: 000101
	// n: 34: 100010
	// n: 03: 000011
	// n: 30: 011110

	fmt.Printf("numbers: %v\n", numbers)

	return strings.Join(numbers, "")
}
