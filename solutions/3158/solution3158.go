package solution3158

// ============================================================================
// 3158. Find the XOR of Numbers Which Appear Twice
// URL: https://leetcode.com/problems/find-the-xor-of-numbers-which-appear-twice/description/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3158---Find-the-XOR-of-Numbers-Which-Appear-Twice
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_duplicateNumbersXOR-24    	197307122	         6.008 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func duplicateNumbersXOR(nums []int) int {
	result := 0
	freq := make([]int, 51)
	for _, n := range nums {
		freq[n]++
		if freq[n] >= 2 {
			result ^= n
		}
	}

	return result
}
