package solution2428

func maxSum(grid [][]int) int {

	const (
		d int = 2
	)

	var m int
	for y := 0; y < len(grid)-d; y++ {
		for x := 0; x < len(grid[y])-d; x++ {
			n := 0
			n += grid[y][x]
			n += grid[y][x+1]
			n += grid[y][x+2]
			n += grid[y+1][x+1]
			n += grid[y+2][x]
			n += grid[y+2][x+1]
			n += grid[y+2][x+2]
			m = max(m, n)
		}
	}

	return m
}
