package solution3274

// ============================================================================
// 3274. Check if Two Chessboard Squares Have the Same Color
// URL: https://leetcode.com/problems/check-if-two-chessboard-squares-have-the-same-color/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3274---Check-If-Two-Chessboard-Squares-Have-the-Same-Color
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_checkTwoChessboards
	Benchmark_checkTwoChessboards-24    	1000000000	         0.8116 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func checkTwoChessboards(c1 string, c2 string) bool {
	return (c1[0]+c2[0]+c1[1]+c2[1])%2 == 0
}
