package solution7

import (
	"math"
	"slices"
	"strconv"
)

// ============================================================================
// 7. Reverse Integer
// URL: https://leetcode.com/problems/reverse-integer/
// ============================================================================

func reverse(x int) int {

	s := strconv.Itoa(x)

	sign := false
	if s[0] == '-' {
		s = s[1:]
		sign = true
	}

	b := []byte(s)

	slices.Reverse(b)

	if sign {
		b = append([]byte{byte(45)}, b...)
	}

	n, _ := strconv.Atoi(string(b))
	if n > math.MinInt32 && n < math.MaxInt32 {
		return n
	}

	return 0
}
