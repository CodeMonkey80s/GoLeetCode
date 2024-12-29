package solution2545

// ============================================================================
// 2545. Sort the Students by Their Kth Score
// URL: https://leetcode.com/problems/sort-the-students-by-their-kth-score/
// ============================================================================

import "slices"

func sortTheStudents(score [][]int, k int) [][]int {
	slices.SortFunc(score, func(a []int, b []int) int {
		return b[k] - a[k]
	})

	return score
}
