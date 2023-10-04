package solution455

import "sort"

// ============================================================================
// 455. Assign Cookies
// URL: https://leetcode.com/problems/assign-cookies/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_findContentChildren-24    	23224447	        56.34 ns/op	      48 B/op	       2 allocs/op
	PASS

*/

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
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
