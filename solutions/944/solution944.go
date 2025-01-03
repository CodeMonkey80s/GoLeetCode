package solution944

import (
	"slices"
)

// ============================================================================
// 944. Delete Columns to Make Sorted
// URL: https://leetcode.com/problems/delete-columns-to-make-sorted/
// ============================================================================

/*

	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/944---Delete-Columns-to-Make-Sorted
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_minDeletionSize
	Benchmark_minDeletionSize-24    	32172034	        34.71 ns/op	      24 B/op	       1 allocs/op
	PASS

*/

func minDeletionSize(strs []string) int {
	ans := 0
	st := make([]int, len(strs))
	for k := 0; k < len(strs[0]); k++ {
		for r := 0; r < len(strs); r++ {
			val := int(strs[r][k] - 96)
			st[r] = val
		}
		ok := slices.IsSorted(st)
		if !ok {
			ans++
		}
	}

	return ans
}
