package solution3028

// ============================================================================
// 3028. Ant on the Boundary
// URL: https://leetcode.com/problems/ant-on-the-boundary/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3028---Ant-on-the-Boundary
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	BenchmarkGetEncryptedString
	BenchmarkGetEncryptedString-24    	484687742	         2.466 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func returnToBoundaryCount(nums []int) int {
	n := 0
	pos := 0
	for _, v := range nums {
		pos += v
		if pos == 0 {
			n++
		}
	}

	return n
}
