package solution2391

// ============================================================================
// 2391. Minimum Amount of Time to Collect Garbage
// URL: https://leetcode.com/problems/minimum-amount-of-time-to-collect-garbage/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2391
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_garbageCollectionV2
	Benchmark_garbageCollectionV2-24    	59838808	        19.83 ns/op	       0 B/op	       0 allocs/op
	Benchmark_garbageCollectionV1
	Benchmark_garbageCollectionV1-24    	60818158	        18.16 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func garbageCollectionV2(garbage []string, travel []int) int {

	for i := 1; i < len(travel); i++ {
		travel[i] = travel[i-1] + travel[i]
	}

	var p int
	var g int
	var m int
	var time int
	for i := len(garbage) - 1; i >= 0; i-- {
		time += len(garbage[i])

		for j := 0; j < len(garbage[i]); j++ {
			if i == 0 {
				continue
			}
			if p == 0 && garbage[i][j] == 'P' {
				time += travel[i-1]
				p = i
			}
			if g == 0 && garbage[i][j] == 'G' {
				time += travel[i-1]
				g = i
			}
			if m == 0 && garbage[i][j] == 'M' {
				time += travel[i-1]
				m = i
			}
		}
	}

	return time
}

func garbageCollectionV1(garbage []string, travel []int) int {

	var g int
	var p int
	var m int

	for i := len(garbage) - 1; i >= 0; i-- {
		for j := 0; j < len(garbage[i]); j++ {
			switch {
			case garbage[i][j] == 'P':
				p += 1
			case garbage[i][j] == 'G':
				g += 1
			case garbage[i][j] == 'M':
				m += 1
			}
		}
		if i == 0 {
			continue
		}
		if p > 0 {
			p += travel[i-1]
		}
		if g > 0 {
			g += travel[i-1]
		}
		if m > 0 {
			m += travel[i-1]
		}
	}

	return g + p + m
}
