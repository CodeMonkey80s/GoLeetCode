package solution48

import "slices"

func rotate(matrix [][]int) {
	for y := 0; y < len(matrix); y++ {
		for x := y + 1; x < len(matrix); x++ {
			matrix[y][x], matrix[x][y] = matrix[x][y], matrix[y][x]
		}
		slices.Reverse(matrix[y])
	}
}
