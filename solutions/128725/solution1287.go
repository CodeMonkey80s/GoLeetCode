package solution1287

// ============================================================================
// 1287. Element Appearing More Than 25% In Sorted Array
// URL: https://leetcode.com/problems/find-if-digit-game-can-be-won/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1287---Element-Appearing-More-Than-25p-In-Sorted-Array
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_findSpecialInteger
	Benchmark_findSpecialInteger-24    	 9925416	       119.4 ns/op	     288 B/op	       1 allocs/op
	PASS
*/

func findSpecialInteger(arr []int) int {
	freq := make(map[int]int, len(arr))

	for _, v := range arr {
		freq[v]++
		if freq[v] > len(arr)/4 {
			return v
		}
	}

	return 0
}
