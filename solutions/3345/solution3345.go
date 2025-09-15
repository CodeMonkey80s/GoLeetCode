package solution3345

import "strconv"

func smallestNumber(n int, t int) int {
	for i := n; i <= n+t; i++ {
		digits := strconv.Itoa(i)
		product := 1
		for j := 0; j < len(digits); j++ {
			product *= int(digits[j] - '0')
		}
		if product%t == 0 {
			return i
		}
	}

	return -1
}
