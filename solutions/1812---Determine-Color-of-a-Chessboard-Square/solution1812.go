package solution1812

// ============================================================================
// 1812. Determine Color of a Chessboard Square
// URL: https://leetcode.com/problems/determine-color-of-a-chessboard-square/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1812---Determine-Color-of-a-Chessboard-Square
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_squareIsWhiteV2
	Benchmark_squareIsWhiteV2-24    	1000000000	         0.5260 ns/op	       0 B/op	       0 allocs/op
	Benchmark_squareIsWhiteV1
	Benchmark_squareIsWhiteV1-24    	1000000000	         0.5249 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func squareIsWhiteV2(coordinates string) bool {
	return (coordinates[0]+coordinates[1])%2 != 0
}

func squareIsWhiteV1(coordinates string) bool {
	r1 := 0b_1010101010
	r2 := 0b_0101010101
	row := 0
	x := byte(coordinates[0]) - 'a'
	y := byte(coordinates[1]) - '0'
	if y%2 == 0 {
		row = r2
	} else {
		row = r1
	}
	return row&(1<<(8-x)) > 0
}
