package solution3516

import "math"

// ============================================================================
// 3515. Find Closest Person
// URL: https://leetcode.com/problems/find-closest-person/
// ============================================================================

func findClosest(x int, y int, z int) int {
	d1 := math.Abs(float64(z - x))
	d2 := math.Abs(float64(z - y))

	switch {
	case d1 < d2:
		return 1
	case d2 < d1:
		return 2
	}

	return 0
}
