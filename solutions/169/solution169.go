package solution169

// ============================================================================
// 169. Majority Element
// URL: https://leetcode.com/problems/majority-element/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/169---Majority-Element
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_majorityElement-24    	60238834	        19.94 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}
	return 0
}
