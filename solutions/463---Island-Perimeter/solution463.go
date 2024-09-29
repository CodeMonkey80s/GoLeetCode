package solution463

// ============================================================================
// 463. Island Perimeter
// URL: https://leetcode.com/problems/island-perimeter/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/463---Island-Perimeter
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_islandPerimeter
	Benchmark_islandPerimeter-24    	56296902	        39.46 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func islandPerimeter(grid [][]int) int {

	var perimeter int

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 0 {
				continue
			}

			sides := 4

			if y > 0 && grid[y-1][x] == 1 {
				sides--
			}
			if x > 0 && grid[y][x-1] == 1 {
				sides--
			}
			if y < len(grid)-1 && grid[y+1][x] == 1 {
				sides--
			}
			if x < len(grid[y])-1 && grid[y][x+1] == 1 {
				sides--
			}

			perimeter += sides
		}
	}

	return perimeter
}
