package solution944

import (
	"sort"
)

// ============================================================================
// 944. Delete Columns to Make Sorted
// URL: https://leetcode.com/problems/delete-columns-to-make-sorted/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_minDeletionSize-24    	15351274	        83.11 ns/op	      96 B/op	       4 allocs/op
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
		ok := sort.IntsAreSorted(st)
		if !ok {
			ans++
		}
	}

	return ans
}
