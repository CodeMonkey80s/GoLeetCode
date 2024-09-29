package solution455

import (
	"slices"
)

// ============================================================================
// 455. Assign Cookies
// URL: https://leetcode.com/problems/assign-cookies/
// ============================================================================

/*

	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/455---Assign-Cookies
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_findContentChildren
	Benchmark_findContentChildren-24    	52557244	        22.56 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}
	slices.Sort(g)
	slices.Sort(s)
	ans := 0
	for j := 0; j < len(s); j++ {
		for i := 0; i < len(g); i++ {
			if s[j] != 0 && g[i] != 0 && s[j] >= g[i] {
				ans++
				g[i] = 0
				s[j] = 0
				break
			}
		}
	}
	return ans
}
