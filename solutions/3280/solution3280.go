package solution3280

import (
	"fmt"
	"slices"
	"strconv"
)

// ============================================================================
// 3280. Convert Date to Binary
// URL: https://leetcode.com/problems/convert-date-to-binary/
// ============================================================================

func convertDateToBinaryV2(date string) string {

	var (
		buf = make([]rune, 0, 24)
		p   = []int{1000, 100, 10, 1, 0, 10, 1, 0, 10, 1, 0}
	)

	numToBin := func(num int, dash bool) {
		b := make([]rune, 0, 4)
		for num > 0 {
			b = append(b, rune(num%2+'0'))
			num = num / 2
		}
		slices.Reverse(b)
		buf = append(buf, b...)
		if dash {
			buf = append(buf, '-')
		}
	}

	var v int
	num := 0
	idx := 0
	for i := 0; i < len(date); i++ {
		if date[i] == '-' {
			numToBin(num, true)
			num = 0
			idx++
			continue
		}
		v = int(date[i]-'0') * p[idx]
		num += v
		idx++
		if i == len(date)-1 {
			numToBin(num, false)
		}
	}

	return string(buf)
}

func convertDateToBinaryV1(date string) string {
	a, _ := strconv.Atoi(date[:4])
	b, _ := strconv.Atoi(date[5:7])
	c, _ := strconv.Atoi(date[8:])
	return fmt.Sprintf("%b-%b-%b", a, b, c)
}
