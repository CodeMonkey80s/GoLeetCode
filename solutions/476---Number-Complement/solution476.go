package solution476

import (
	"strconv"
)

// ============================================================================
// 476. Number Complement
// URL: https://leetcode.com/problems/number-complement/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_findComplement-24    	14815994	        82.36 ns/op	       8 B/op	       3 allocs/op
	Benchmark_findComplement-24    	24780900	        56.66 ns/op	       8 B/op	       3 allocs/op

*/

func findComplement(num int) int {
	n1 := strconv.FormatInt(int64(num), 2)
	n2 := ""
	for i := len(n1) - 1; i >= 0; i-- {
		if n1[i] == '0' {
			n2 = "1" + n2
		} else {
			n2 = "0" + n2
		}
	}
	ans, err := strconv.ParseInt(n2, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(ans)
}
