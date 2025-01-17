package solution2373

// ============================================================================
// 2373. Largest Local Values in a Matrix
// URL: https://leetcode.com/problems/largest-local-values-in-a-matrix/
// ============================================================================

func largestLocal(grid [][]int) [][]int {

	const (
		d int = 2
	)

	output := make([][]int, len(grid)-d)
	for y := 0; y < len(grid)-d; y++ {
		output[y] = make([]int, len(grid)-d)
		for x := 0; x < len(grid[y])-d; x++ {
			n := 0
			n = max(grid[y][x], n)
			n = max(grid[y][x+1], n)
			n = max(grid[y][x+2], n)
			n = max(grid[y+1][x], n)
			n = max(grid[y+1][x+1], n)
			n = max(grid[y+1][x+2], n)
			n = max(grid[y+2][x], n)
			n = max(grid[y+2][x+1], n)
			n = max(grid[y+2][x+2], n)
			output[y][x] = n
		}
	}

	return output
}
