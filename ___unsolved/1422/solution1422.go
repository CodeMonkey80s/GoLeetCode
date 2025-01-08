package solution1422

import (
	"fmt"
	"math"
)

func maxScore(s string) int {

	sum := make([]int, len(s))

	//     0 1 1 1 0 1
	// 0:  1 1 1 1 2 2
	// 1:  0 1 2 3 3 4
	// s:  1         4 = 5

	for i := 0; i < len(s); i++ {
		sum[i] = int(s[i] - '0')
		if i > 0 {
			sum[i] += sum[i-1]
		}
	}

	m := math.MinInt
	for i := 0; i <= len(s)-2; i++ {
		// c := 0
	}

	fmt.Printf("sum: %v\n", sum)

	return m
}
