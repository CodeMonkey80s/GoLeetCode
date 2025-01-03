package solution1476

// ============================================================================
// 1476. Subrectangle Queries
// URL: https://leetcode.com/problems/subrectangle-queries/description/
// ============================================================================

type SubrectangleQueries struct {
	grid [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{grid: rectangle}
}

func (s *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for y := row1; y <= row2; y++ {
		for x := col1; x <= col2; x++ {
			s.grid[y][x] = newValue
		}
	}
}

func (s *SubrectangleQueries) GetValue(row int, col int) int {
	return s.grid[row][col]
}
