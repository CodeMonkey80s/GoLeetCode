package solution3750

import "strconv"

func minimumFlips(n int) int {
	bits := strconv.FormatInt(int64(n), 2)
	var flips int
	for i := 0; i < len(bits); i++ {
		a := i
		b := len(bits) - 1 - i
		if bits[a] != bits[b] {
			flips++
		}
	}
	return flips
}
