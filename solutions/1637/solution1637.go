package solution1637

// ============================================================================
// 1637. Widest Vertical Area Between Two Points Containing No Points
// URL: https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points/
// ============================================================================

import (
	"slices"
)

func maxWidthOfVerticalArea(points [][]int) int {

	slices.SortFunc(points, func(a []int, b []int) int {
		return a[0] - b[0]
	})

	var m int
	for i := 0; i < len(points)-1; i++ {
		diff := points[i+1][0] - points[i][0]
		m = max(m, diff)
	}

	return m
}
