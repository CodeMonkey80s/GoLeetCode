package solution2482

// ============================================================================
// 2482. Difference Between Ones and Zeros in Row and Column
// URL: https://leetcode.com/problems/difference-between-ones-and-zeros-in-row-and-column/
// ============================================================================

func onesMinusZeros(grid [][]int) [][]int {

	onesRow := make([]int, len(grid))
	onesCol := make([]int, len(grid[0]))
	zerosRow := make([]int, len(grid))
	zerosCol := make([]int, len(grid[0]))

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			n := grid[y][x]
			if n == 0 {
				zerosRow[y]++
				zerosCol[x]++
			} else {
				onesRow[y]++
				onesCol[x]++
			}
		}
	}

	output := make([][]int, 0, len(grid))
	for y := 0; y < len(grid); y++ {
		row := make([]int, len(grid[0]))
		for x := 0; x < len(grid[0]); x++ {
			row[x] = onesRow[y] + onesCol[x] - zerosRow[y] - zerosCol[x]
		}
		output = append(output, row)
	}

	return output
}
