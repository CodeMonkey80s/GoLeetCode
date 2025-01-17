package solution2037

import "slices"

// ============================================================================
// 2037. Minimum Number of Moves To Seat Everyone
// URL: https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/
// ============================================================================

func minMovesToSeat(seats []int, students []int) int {

	slices.Sort(seats)
	slices.Sort(students)

	var moves int
	for i := range seats {
		d := students[i] - seats[i]
		if d < 0 {
			d = -d
		}
		moves += d
	}

	return moves
}
