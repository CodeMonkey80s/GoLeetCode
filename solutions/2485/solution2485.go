package solution2485

// ============================================================================
// 2485. Find the Pivot Integer
// URL: https://leetcode.com/problems/find-the-pivot-integer/
// ============================================================================

func pivotInteger(n int) int {

	sum1 := make([]int, n+1)
	sum2 := make([]int, n+2)

	for i := 1; i <= n; i++ {
		sum1[i] = sum1[i-1] + i
	}

	for i := n; i >= 1; i-- {
		sum2[i] = sum2[i+1] + i
	}

	sum1 = sum1[1:]
	sum2 = sum2[1:]
	sum2 = sum2[:len(sum2)-1]

	for i := 0; i < n; i++ {
		if sum1[i] == sum2[i] {
			return i + 1
		}
	}

	return -1
}
