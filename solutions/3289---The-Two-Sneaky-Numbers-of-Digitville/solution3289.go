package solution3289

import (
	"slices"
)

// ============================================================================
// 3289. The Two Sneaky Numbers of Digitville
// URL: https://leetcode.com/problems/the-two-sneaky-numbers-of-digitville/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3289---The-Two-Sneaky-Numbers-of-Digitville
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_getSneakyNumbers
	Benchmark_getSneakyNumbers-24    	11988784	       105.0 ns/op	      96 B/op	       2 allocs/op
	PASS
*/

func getSneakyNumbers(nums []int) []int {
	ans := make([]int, len(nums))
	freq := make(map[int]int, len(nums))
	for _, n := range nums {
		freq[n]++
		v, ok := freq[n]
		if ok {
			if v == 2 {
				ans = append(ans, n)
			}
		}
	}

	slices.Sort(ans)

	return ans
}
