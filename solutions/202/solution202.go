package solution202

import (
	"math"
	"strconv"
)

func isHappyV2(n int) bool {

	sum := func(n int) int {
		var sum int
		for n > 0 {
			d := n % 10
			sum += d * d
			n /= 10
		}

		return sum
	}

	seen := make(map[int]struct{})

	for {
		n = sum(n)
		switch {
		case n == 1:
			return true
		case func() bool { _, ok := seen[n]; return ok }():
			return false
		default:
			seen[n] = struct{}{}
		}
	}
}

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
