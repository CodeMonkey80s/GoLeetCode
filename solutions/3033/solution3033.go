package solution3033

func modifiedMatrix(matrix [][]int) [][]int {

	for x := 0; x < len(matrix[0]); x++ {

		m := -1
		found := false
		for y := 0; y < len(matrix); y++ {
			m = max(m, matrix[y][x])
			if matrix[y][x] == -1 {
				found = true
			}
		}

		if !found {
			continue
		}

		for y := 0; y < len(matrix); y++ {
			if matrix[y][x] == -1 {
				matrix[y][x] = m
			}
		}
	}

	return matrix
}
