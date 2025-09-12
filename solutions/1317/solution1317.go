package solution1317

import "strconv"

func getNoZeroIntegers(n int) []int {
	noZero := func(n int) bool {
		s := strconv.Itoa(n)
		for _, ch := range s {
			if ch == '0' {
				return false
			}
		}
		return true
	}

	a := n - 1
	b := n - a

	for {
		if noZero(a) && noZero(b) {
			return []int{a, b}
		}
		a--
		b++
	}
}
