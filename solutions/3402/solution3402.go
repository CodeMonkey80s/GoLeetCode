package solution3402

// ============================================================================
// 3402. Minimum Operations to Make Columns Strictly Increasing
// URL: https://leetcode.com/problems/minimum-operations-to-make-columns-strictly-increasing/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3402
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_minimumOperations
	Benchmark_minimumOperations-24    	163613364	        11.40 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func minimumOperations(grid [][]int) int {

	var d int
	var ops int
	for x := 0; x < len(grid[0]); x++ {
		for y := 1; y <= len(grid)-1; y++ {
			if grid[y][x] <= grid[y-1][x] {
				d = grid[y-1][x] - grid[y][x] + 1
				grid[y][x] += d
				ops += d
			}
		}
	}

	return ops
}
