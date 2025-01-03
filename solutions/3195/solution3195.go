package solution3195

// ============================================================================
// 3195. Find the Minimum Area to Cover All Ones I
// URL: https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-i/
// ============================================================================

func minimumArea(grid [][]int) int {
	minx := len(grid[0])
	miny := len(grid[0])
	maxx := 0
	maxy := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 0 {
				continue
			}
			minx = min(minx, x)
			maxx = max(maxx, x)
			miny = min(miny, y)
			maxy = max(maxy, y)
		}
	}

	return (maxy + 1 - miny) * (maxx + 1 - minx)
}
