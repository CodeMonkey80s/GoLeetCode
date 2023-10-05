package solution1436

// ============================================================================
// 1436. Destination City
// URL: https://leetcode.com/problems/destination-city/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1436---Destination-City
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_destCity-24    	33474986	        35.79 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func destCity(paths [][]string) string {
	m := make(map[string]int)
	for _, v := range paths {
		src := v[0]
		dst := v[1]
		_, ok := m[dst]
		if !ok {
			m[src] = 1
		} else {
			m[src]++
		}
	}
	for _, v := range paths {
		dst := v[1]
		_, ok := m[dst]
		if !ok {
			return dst
		}
	}
	return ""
}
