package solution2500

import "slices"

func deleteGreatestValue(grid [][]int) int {

	for i := 0; i < len(grid); i++ {
		slices.Sort(grid[i])
	}

	var m int
	var answer int
	for x := len(grid[0]) - 1; x >= 0; x-- {
		m = grid[0][x]
		for y := 0; y <= len(grid)-1; y++ {
			m = max(m, grid[y][x])
			if y > len(grid)-1 {
				m = max(m, grid[y+1][x])
			}
		}
		answer += m
	}

	return answer
}
