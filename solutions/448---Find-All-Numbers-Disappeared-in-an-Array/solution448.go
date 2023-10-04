package solution448

// ============================================================================
// 448. Find All Numbers Disappeared in an Array
// URL: https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_find-24    	 7476871	       156.5 ns/op	      24 B/op	       2 allocs/op

*/

func findDisappearedNumbers(nums []int) []int {
	ans := make([]int, 0)
	m := make(map[int]int)
	v := 0
	maxval := len(nums)
	for i := 0; i < maxval; i++ {
		v = nums[i]
		_, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}
	for i := 1; i <= maxval; i++ {
		_, ok := m[i]
		if !ok {
			ans = append(ans, i)
		}
	}
	return ans
}
