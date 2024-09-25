package solution219

// ============================================================================
// 219. Contains Duplicate II
// URL: https://leetcode.com/problems/contains-duplicate-ii/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/219---Contains-Duplicate-II
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_containsNearbyDuplicate
	Benchmark_containsNearbyDuplicate-24    	31926517	        31.42 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	track := make(map[int]int, len(nums))
	for i, v := range nums {
		index, ok := track[v]
		if ok && (i-index) <= k {
			return true
		}
		track[v] = i
	}
	return false

}
