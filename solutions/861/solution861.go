package solution861

// ============================================================================
// 861. Score After Flipping Matrix
// URL: https://leetcode.com/problems/score-after-flipping-matrix/
// ============================================================================

import (
	"strconv"
	"strings"
)

func matrixScore(grid [][]int) int {

	for y := 0; y < len(grid); y++ {
		if grid[y][0] == 0 {
			for x := 0; x < len(grid[0]); x++ {
				if grid[y][x] == 0 {
					grid[y][x] = 1
				} else {
					grid[y][x] = 0
				}
			}
		}
	}

	for x := 0; x < len(grid[0]); x++ {
		c0 := 0
		c1 := 0
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == 0 {
				c0++
			} else {
				c1++
			}
		}
		if c0 > c1 {
			for y := 0; y < len(grid); y++ {
				if grid[y][x] == 0 {
					grid[y][x] = 1
				} else {
					grid[y][x] = 0
				}
			}
		}
	}

	var sb strings.Builder

	score := 0
	for y := 0; y < len(grid); y++ {
		sb.Grow(len(grid[0]))
		for x := 0; x < len(grid[0]); x++ {
			sb.WriteString(strconv.Itoa(grid[y][x]))
		}
		n, _ := strconv.ParseInt(sb.String(), 2, 64)
		score += int(n)
		sb.Reset()
	}

	return score
}
