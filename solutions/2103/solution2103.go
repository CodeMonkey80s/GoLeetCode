package solution2103

// ============================================================================
// 2103. Rings and Rods
// URL: https://leetcode.com/problems/rings-and-rods/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2103---Rigns-and-Rods
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countPoints-24    	100000000	        11.23 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func countPoints(rings string) int {
	type Rod struct {
		r bool
		g bool
		b bool
	}
	poles := make([]Rod, 10)
	ans := 0
	for i := 0; i < len(rings); i += 2 {
		color := rings[i]
		idx := byte(rings[i+1]) - '0'
		switch color {
		case 'R':
			poles[idx].r = true
		case 'G':
			poles[idx].g = true
		case 'B':
			poles[idx].b = true
		}
	}
	for i := range poles {
		if poles[i].r && poles[i].g && poles[i].b {
			ans++
		}
	}
	return ans
}
