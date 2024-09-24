package solution26

import (
	"slices"
)

// ============================================================================
// 26. Remove Duplicates from Sorted Array
// URL: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// ============================================================================

func removeDuplicates_stdlib(nums []int) int {
	nums = slices.Compact(nums)
	return len(nums)
}

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/26---Remove-Duplicates-from-Sorted-Array
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_removeDuplicates
	Benchmark_removeDuplicates-24    	16625343	        66.07 ns/op	      24 B/op	       2 allocs/op
	PASS
*/

func removeDuplicates(nums []int) int {
	m := make(map[int]int, len(nums))
	s := make([]int, 0, len(nums))
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n]++
			s = append(s, n)
		}
	}
	copy(nums, s)
	l2 := len(s)
	nums = nums[0:l2:l2]
	return len(nums)
}
