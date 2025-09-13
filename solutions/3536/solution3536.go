package solution3536

import (
	"slices"
	"strconv"
)

func maxProduct(n int) int {

	s := []byte(strconv.Itoa(n))
	slices.SortFunc(s, func(i, j byte) int {
		if i > j {
			return -1
		}
		return 0
	})

	return int((s[0] - '0') * (s[1] - '0'))
}
