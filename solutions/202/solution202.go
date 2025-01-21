package solution202

import (
	"math"
	"strconv"
)

func isHappy(n int) bool {

	hash := make(map[int]int)
	for {

		if n == 1 {
			return true
		}

		numbers := strconv.Itoa(n)

		var sum int
		for _, digit := range numbers {
			v := int(math.Pow(float64(digit-'0'), 2))
			sum += v
		}

		n = sum
		hash[n]++
		c, ok := hash[n]
		if ok && c >= 2 {
			break
		}
	}

	return false
}
