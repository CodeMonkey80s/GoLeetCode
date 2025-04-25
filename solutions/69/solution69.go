package solution69

// ============================================================================
// 69. Sqrt(x)
// https://leetcode.com/problems/sqrtx/
// ============================================================================

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	var i = 1
	var result = 1
	for result <= x {
		i++
		result = i * i
	}
	return i - 1
}
