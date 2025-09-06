package solution3668

// ============================================================================
// 3668. Restore Finish Order
// URL: https://leetcode.com/problems/restore-finishing-order/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3668
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_gcdOfOddEvenSums
	Benchmark_gcdOfOddEvenSums-24    	35322500	        31.98 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func recoverOrder(order []int, friends []int) []int {
	freq := make(map[int]struct{})
	for i := 0; i < len(friends); i++ {
		freq[friends[i]] = struct{}{}
	}

	var idx int
	for i := 0; i < len(order); i++ {
		if _, ok := freq[order[i]]; ok {
			friends[idx] = order[i]
			idx++
		}
	}

	return friends
}
