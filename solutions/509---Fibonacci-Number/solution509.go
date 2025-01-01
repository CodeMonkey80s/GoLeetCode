package solution509

// ============================================================================
// 509. Fibonacci Number
// URL: https://leetcode.com/problems/fibonacci-number/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/509---Fibonacci-Number
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_fib
	Benchmark_fib-24    	88254508	        16.30 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

var cache = map[int]int{1: 1}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	v1, ok1 := cache[n-1]
	v2, ok2 := cache[n-2]

	if ok2 && ok1 {
		return v1 + v2
	} else if !ok2 && ok1 {
		cache[n] = v1 + fib(n-2)
	} else if ok2 {
		cache[n] = fib(n-1) + v2
	} else {
		cache[n] = fib(n-1) + fib(n-2)
	}

	return cache[n]
}
