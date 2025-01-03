package solution1331

import (
	"slices"
)

// ============================================================================
// 1331
// URL: https://leetcode.com/problems/binary-number-with-alternating-bits/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1331---Rank-Transform-of-an-Array
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_arrayRankTransform
	Benchmark_arrayRankTransform-24    	13334482	        85.24 ns/op	      32 B/op	       1 allocs/op
	PASS
*/

func arrayRankTransform(arr []int) []int {

	buf := slices.Clone(arr)

	slices.Sort(buf)

	m := make(map[int]int)

	rank := 1
	for i := range buf {
		if i >= 1 && buf[i] == buf[i-1] {
			rank--
		}
		m[buf[i]] = rank
		rank++
	}

	for i := range arr {
		arr[i] = m[arr[i]]
	}

	return arr
}
