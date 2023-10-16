package solution412

import "strconv"

// ============================================================================
// 412. Fizz Buzz
// URL: https://leetcode.com/problems/fizz-buzz/
// ============================================================================

func fizzBuzz(n int) []string {
	sl := make([]string, 0)
	var i int
	var s string
	for j := 0; j < n; j++ {
		i = j + 1
		switch {
		case i%3 == 0 && i%5 == 0:
			s = "FizzBuzz"
		case i%3 == 0:
			s = "Fizz"
		case i%5 == 0:
			s = "Buzz"
		default:
			s = strconv.Itoa(i)
		}
		sl = append(sl, s)
	}
	return sl
}
