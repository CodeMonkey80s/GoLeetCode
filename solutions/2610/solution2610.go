package solution2610

// ============================================================================
// 2610. Convert an Array Into a 2D Array With Conditions
// URL: https://leetcode.com/problems/convert-an-array-into-a-2d-array-with-conditions/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2610
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_findMatrix
	Benchmark_findMatrix-24    	 6725511	       182.9 ns/op	     264 B/op	       7 allocs/op
	PASS
*/

import (
	"slices"
)

func findMatrix(nums []int) [][]int {

	output := make([][]int, len(nums))
	for _, num := range nums {

		if len(output) == 0 {
			output[0] = append(output[0], num)
		}

		idx, ok := find(output, num)
		if ok {
			output[idx] = append(output[idx], num)
			continue
		}

		output[len(output)] = append(output[len(output)], num)

	}

	for i := range output {
		if len(output[i]) == 0 {
			output = output[:i]
			break
		}
	}

	output = slices.Clip(output)

	return output
}

func find(output [][]int, num int) (int, bool) {
	idx := 0
	for i, row := range output {
		if !slices.Contains(row, num) {
			return i, true
		}
	}

	return idx, false
}
