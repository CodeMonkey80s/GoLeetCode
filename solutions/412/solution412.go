package solution412

import "strconv"

// ============================================================================
// 412. Fizz Buzz
// URL: https://leetcode.com/problems/fizz-buzz/
// ============================================================================

/*
	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/412---Fizz-Buzz
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_fizzBuzz-24    	43119661	        28.05 ns/op	      48 B/op	       1 allocs/op
	PASS

*/

func fizzBuzz(n int) []string {
	sl := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			sl = append(sl, "FizzBuzz")
		case i%3 == 0:
			sl = append(sl, "Fizz")
		case i%5 == 0:
			sl = append(sl, "Buzz")
		default:
			sl = append(sl, strconv.Itoa(i))
		}
	}
	return sl
}
