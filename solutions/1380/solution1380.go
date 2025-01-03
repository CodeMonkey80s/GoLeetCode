package solution1380

// ============================================================================
// 1380. Lucky Numbers in a Matrix
// URL: https://leetcode.com/problems/lucky-numbers-in-a-matrix/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1380
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_luckyNumbers
	Benchmark_luckyNumbers-24    	 3312883	       328.7 ns/op	     104 B/op	       5 allocs/op
	PASS
*/

import (
	"math"
)

func luckyNumbers(matrix [][]int) []int {

	rows := make(map[int]bool)
	cols := make(map[int]bool)

	for y := 0; y < len(matrix); y++ {
		minNumber := math.MaxInt
		for x := 0; x < len(matrix[y]); x++ {
			minNumber = min(minNumber, matrix[y][x])
		}
		rows[minNumber] = true
	}

	for x := 0; x < len(matrix[0]); x++ {
		maxNumber := math.MinInt
		for y := 0; y < len(matrix); y++ {
			maxNumber = max(maxNumber, matrix[y][x])
		}
		cols[maxNumber] = true
	}

	for rowVal := range rows {
		for colVal := range cols {
			if rowVal == colVal {
				return []int{rowVal}
			}
		}
	}

	return []int{}
}
