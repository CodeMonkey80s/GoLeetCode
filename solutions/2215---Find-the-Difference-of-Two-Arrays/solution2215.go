package solution2215

import (
	"slices"
)

/*
goos: linux
goarch: amd64
pkg: GoLeetCode/solutions/2215---Find-the-Difference-of-Two-Arrays
cpu: 13th Gen Intel(R) Core(TM) i7-13700K
Benchmark_findDifference
Benchmark_findDifference-24    	 6763864	       177.0 ns/op	      96 B/op	       3 allocs/op
PASS
*/
func findDifference(nums1 []int, nums2 []int) [][]int {

	freq1 := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		freq1[num]++
	}

	freq2 := make(map[int]int, len(nums2))
	for _, num := range nums2 {
		freq2[num]++
	}

	dist1 := make([]int, 0, len(nums1))
	dist2 := make([]int, 0, len(nums2))

	for _, num := range nums1 {
		if slices.Contains(dist1, num) {
			continue
		}
		_, ok := freq2[num]
		if !ok {
			dist1 = append(dist1, num)
		}
	}

	for _, num := range nums2 {
		if slices.Contains(dist2, num) {
			continue
		}
		_, ok := freq1[num]
		if !ok {
			dist2 = append(dist2, num)
		}
	}

	return [][]int{dist1, dist2}
}
