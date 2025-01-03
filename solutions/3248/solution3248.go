package solution3248

// ============================================================================
// 3248. Snake in Matrix
// URL: https://leetcode.com/problems/snake-in-matrix/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3248---Snake-in-Matrix
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_finalPositionOfSnake
	Benchmark_finalPositionOfSnake-24    	300523683	         4.002 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func finalPositionOfSnake(n int, commands []string) int {
	var x, y int
	for _, cmd := range commands {
		switch cmd {
		case "UP":
			y--
		case "DOWN":
			y++
		case "RIGHT":
			x++
		case "LEFT":
			x--
		}
	}

	return y*n + x
}
