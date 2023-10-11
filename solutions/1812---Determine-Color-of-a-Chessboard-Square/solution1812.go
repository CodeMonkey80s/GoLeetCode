package solution1812

// ============================================================================
// 1812. Determine Color of a Chessboard Square
// URL: https://leetcode.com/problems/determine-color-of-a-chessboard-square/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1812---Determine-Color-of-a-Chessboard-Square
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_squareIsWhite-24    	1000000000	         0.5246 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func squareIsWhite(coordinates string) bool {
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
	val := row & (1 << (8 - x))
	if val > 0 {
		return true
	}
	return false
}
