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
	Benchmark_getSneakyNumbersV2
	Benchmark_getSneakyNumbersV2-24    	203654988	         9.014 ns/op	       0 B/op	       0 allocs/op
	Benchmark_getSneakyNumbersV1
	Benchmark_getSneakyNumbersV1-24    	11519113	       111.7 ns/op	      96 B/op	       2 allocs/op
	PASS
*/

func getSneakyNumbersV2(nums []int) []int {
	var a, b, c int
	freq := make([]int, 101)
	for _, n := range nums {
		freq[n]++
		switch {
		case c == 0 && freq[n] == 2:
			a = n
			c++
		case c == 1 && freq[n] == 2:
			b = n
			if a > b {
				a, b = b, a
			}
		}
	}

	return []int{a, b}
}

func getSneakyNumbersV1(nums []int) []int {
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
