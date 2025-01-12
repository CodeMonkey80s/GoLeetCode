package solution2373

import "fmt"

// ============================================================================
// 2373. Largest Local Values in a Matrix
// URL: https://leetcode.com/problems/largest-local-values-in-a-matrix/
// ============================================================================

func largestLocal(grid [][]int) [][]int {

	n := len(grid)

	ans := make([][]int, n-2)

	sq := func(sx, sy int) [][]int {
		square := make([][]int, n-2)
		for y := sx; y < n-2; y++ {
			row := make([]int, n-2)
			for x := sy; x < n-2; x++ {
				v := grid[sx+x][sy+y]
				row[x] = v
				fmt.Printf("x=%d, y=%d, v=%d\n", x, y, v)
			}
			square = append(square, row)
		}
		return square
	}

	for x := 0; x < n-2; x++ {
		s := sq(x, 0)
		fmt.Printf("s: %v\n", s)
	}

	fmt.Printf("ans: %v\n", ans)

	return ans
}
