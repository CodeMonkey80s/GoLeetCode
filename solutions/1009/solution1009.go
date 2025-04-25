package solution1009

import (
	"fmt"
	"strconv"
)

func bitwiseComplement(n int) int {

	switch {
	case n < 0:
		return 0
	case n == 0:
		return 1
	}

	num := []byte(fmt.Sprintf("%b", n))
	for i := 0; i < len(num); i++ {
		switch num[i] {
		case '0':
			num[i] = '1'
		case '1':
			num[i] = '0'
		}
	}

	val, _ := strconv.ParseInt(string(num), 2, 64)

	return int(val)
}
