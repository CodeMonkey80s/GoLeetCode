package solution1089

// ============================================================================
// 1089. Duplicate Zeros
// URL: https://leetcode.com/problems/duplicate-zeros/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1089---Duplicate-Zeros
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_duplicateZeros-24    	131173022	        15.51 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			copy(arr[i+1:], arr[i:])
			arr[i+1] = 0
			i += 1
		}
	}
}
