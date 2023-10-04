package solution1496

// ============================================================================
// 1496. Path Crossing
// URL: https://leetcode.com/problems/path-crossing/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_isPathCrossing-24    	21440473	        92.86 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func isPathCrossing(path string) bool {
	type Point struct {
		x, y int
	}
	x, y := 0, 0
	m := make(map[Point]int)
	p := Point{
		x: x,
		y: y,
	}
	m[p] = 1
	for _, v := range path {
		switch v {
		case 'N':
			y--
		case 'S':
			y++
		case 'W':
			x--
		case 'E':
			x++
		}
		p = Point{
			x: x,
			y: y,
		}
		_, ok := m[p]
		if ok {
			return true
		} else {
			m[p] = 1
		}
	}
	return false
}
