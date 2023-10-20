package solution1351

// ============================================================================
// 1351. Count Negative Numbers in a Sorted Matrix
// URL: https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/
// ============================================================================

func countNegatives(grid [][]int) int {
	ans := 0
outer:
	for row := 0; row < len(grid); row++ {
		for i := len(grid[row]) - 1; i >= 0; i-- {
			if grid[row][i] < 0 {
				ans++
				continue
			}
			continue outer
		}
	}
	return ans
}
