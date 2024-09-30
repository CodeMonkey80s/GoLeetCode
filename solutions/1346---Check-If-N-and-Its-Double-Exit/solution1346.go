package solution1346

import "slices"

// ============================================================================
// 1346. Check If N and Its Double Exist
// URL: https://leetcode.com/problems/check-if-n-and-its-double-exist/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1346---Check-If-N-and-Its-Double-Exit
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	BenchmarkCheckIfExist
	BenchmarkCheckIfExist-24    	100000000	        10.62 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func checkIfExist(arr []int) bool {
	for i, n := range arr {
		idx := slices.Index(arr, n*2)
		if idx != -1 && idx != i {
			return true
		}
	}

	return false
}
