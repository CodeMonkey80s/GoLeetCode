package solution70

// ============================================================================
// 70. Climbing Stairs
// URL: https://leetcode.com/problems/climbing-stairs/description/
// ============================================================================

func climbStairs(n int) int {
	a, b := 0, 1
	var v = 0
	for i := 0; i <= n; i++ {
		a, b = b, a+b
		v = a
	}
	return v
}
