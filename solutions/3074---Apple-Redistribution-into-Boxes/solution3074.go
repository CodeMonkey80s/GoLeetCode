package solution3074

import (
	"slices"
)

// ============================================================================
// 3074. Apple Redistribution into Boxes
// URL: https://leetcode.com/problems/apple-redistribution-into-boxes/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3074---Apple-Redistribution-into-Boxes
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_minimumBoxes
	Benchmark_minimumBoxes-24    	124226916	        14.35 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func minimumBoxes(apple []int, capacity []int) int {

	slices.Sort(capacity)

	apples := 0
	for i := 0; i < len(apple); i++ {
		apples += apple[i]
	}

	i := len(capacity) - 1
	boxes := 0
	for i >= 0 && apples > 0 {
		apples -= capacity[i]
		boxes++
		i--
	}

	return boxes
}
