package solution830

import (
	"sort"
)

// ============================================================================
// 830. Positions of Large Groups
// URL: https://leetcode.com/problems/positions-of-large-groups/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_largeGroup-24    	 2861828	       464.2 ns/op	     732 B/op	      22 allocs/op
	PASS

*/

func largeGroupPositions(s string) [][]int {
	type item struct {
		s string
		l int
		a int
		b int
	}
	sl := make([][]int, 0)
	grp := ""
	a := 0
	b := 0
	for i := 0; i < len(s); i++ {
		if len(grp) == 0 {
			a = i
		}
		if len(grp) > 0 {
			if grp[0] != s[i] {
				b = i - 1
				if b == 0 {
					b = 1
				}
				if len(grp) >= 3 {
					sl = append(sl, []int{a, b})
				}
				grp = ""
				a = i
			}
		}
		grp += string(s[i])
		if len(grp) == len(s) {
			a = 0
			b = len(s) - 1
			if len(grp) >= 3 {
				sl = append(sl, []int{a, b})
			}
			break
		}
		if i == len(s)-1 {
			if len(grp) >= 3 {
				sl = append(sl, []int{a, a + len(grp) - 1})
			}
		}
	}
	sort.Slice(sl, func(i, j int) bool {
		return sl[i][0] < sl[j][0]
	})
	return sl
}
