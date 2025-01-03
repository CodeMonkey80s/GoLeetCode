package solution2194

// ============================================================================
// 2194. Cells in a Range on an Excel Sheet
// URL: https://leetcode.com/problems/cells-in-a-range-on-an-excel-sheet/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2194---Cells-in-a-Range-on-an-Excell-Sheet
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_cellsInRange-24           	13339546	       111.5 ns/op	      72 B/op	       5 allocs/op
	Benchmark_cellsInRange_string-24    	 6490401	       167.7 ns/op	     120 B/op	       7 allocs/op
	PASS

*/

func cellsInRange(s string) []string {
	minCol := int(s[0] - 'A')
	minRow := int(s[1] - '0')
	maxCol := int(s[3] - 'A')
	maxRow := int(s[4] - '0')
	ans := make([]string, (maxCol-minCol+1)*(maxRow-minRow+1))
	for y := 0; y <= maxCol-minCol; y++ {
		for x := 0; x <= maxRow-minRow; x++ {
			idx := (y * (maxRow - minRow + 1)) + x
			ans[idx] = string(byte(minCol+y)+'A') + string(byte(minRow+x)+'0')
		}
	}
	return ans
}

func cellsInRange_string(s string) []string {
	ans := []string{}
	minCol := s[0] - 'A'
	minRow := s[1] - '0'
	maxCol := s[3] - 'A'
	maxRow := s[4] - '0'
	for y := minCol; y <= maxCol; y++ {
		for x := minRow; x <= maxRow; x++ {
			cell := string(byte(y)+'A') + string(byte(x)+'0')
			ans = append(ans, cell)
		}
	}
	return ans
}
