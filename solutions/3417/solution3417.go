package solution3417

func zigzagTraversal(grid [][]int) []int {

	var output []int
	for y := 0; y < len(grid); y++ {
		x := 0
		d := 1
		if y%2 != 0 {
			x = len(grid[y]) - 1
			d = -1
		}
		for x >= 0 && x < len(grid[y]) {
			switch {
			case y%2 == 0 && x%2 == 0:
				output = append(output, grid[y][x])
			case y%2 != 0 && x%2 != 0:
				output = append(output, grid[y][x])
			}
			x += d
		}
	}

	return output
}
