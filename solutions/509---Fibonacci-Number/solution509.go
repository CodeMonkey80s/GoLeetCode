package solution509

// ============================================================================
// 509. Fibonacci Number
// URL: https://leetcode.com/problems/fibonacci-number/
// ============================================================================

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
