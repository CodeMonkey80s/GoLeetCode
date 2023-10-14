package solution1

// ============================================================================
// 1. Two Sum
// URL: https://leetcode.com/problems/two-sum/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1---Two-Sum
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_twoSum_Iterations-24    	100000000	        13.31 ns/op	      16 B/op	       1 allocs/op
	Benchmark_twoSum_Map-24           	38308172	        29.08 ns/op	      16 B/op	       1 allocs/op
	PASS

*/

func twoSum_Iterations(nums []int, target int) []int {
	maxval := len(nums)
	if maxval < 2 || maxval > 10_000 {
		return []int{}
	}
	if maxval == 2 && nums[0]+nums[1] == target {
		return []int{0, 1}
	}
	for i := 0; i < maxval; i++ {
		for j := i + 1; j < maxval; j++ {
			t := nums[i] + nums[j]
			if t == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// ============================================================================

func twoSum_Map(nums []int, target int) []int {
	maxval := len(nums)
	if maxval < 2 || maxval > 10_000 {
		return []int{}
	}
	if maxval == 2 && nums[0]+nums[1] == target {
		return []int{0, 1}
	}
	var t int
	m := make(map[int]int, maxval)
	for k, v := range nums {
		t = target - v
		_, ok := m[t]
		if ok {
			return []int{m[t], k}
		}
		m[v] = k
	}
	return []int{}
}
