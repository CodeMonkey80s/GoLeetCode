package solution867

func transpose(matrix [][]int) [][]int {

	rows := len(matrix)
	cols := len(matrix[0])
	output := make([][]int, cols)
	for i := 0; i < len(matrix[0]); i++ {
		output[i] = make([]int, rows)
		for j := 0; j < len(matrix); j++ {
			output[i][j] = matrix[j][i]
		}
	}

	return output
}
