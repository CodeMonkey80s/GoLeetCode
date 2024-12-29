package solution1572

// ============================================================================
// 1572. Matrix Diagonal Sum
// URL: https://leetcode.com/problems/matrix-diagonal-sum/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1572
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_diagonalSum
	Benchmark_diagonalSum-24    	504871434	         3.792 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func diagonalSum(mat [][]int) int {

	a := 0
	b := len(mat[0]) - 1
	sum := 0
	for y := 0; y < len(mat); y++ {
		sa := mat[y][a]
		sb := mat[y][b]
		if a != b {
			sum += sa + sb
		} else {
			sum += sa
		}
		a++
		b--
	}

	return sum
}
