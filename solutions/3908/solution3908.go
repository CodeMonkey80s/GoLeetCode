package solution3908

import "strconv"

func validDigit(n int, x int) bool {
	digits := strconv.Itoa(n)
	if digits[0] == byte(x+'0') {
		return false
	}
	for i := 1; i < len(digits); i++ {
		if digits[i] == byte(x+'0') {
			return true
		}
	}
	return false
}
