package solution1399

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1399
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countLargestGroup
	Benchmark_countLargestGroup-24    	31597207	        37.19 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func countLargestGroup(n int) int {
	var m int
	var ints [37]int
	for i := 1; i <= n; i++ {
		c := 0
		num := i
		for num > 0 {
			c += num % 10
			num /= 10
		}
		ints[c]++
		if ints[c] > m {
			m = ints[c]
		}
	}

	var largest int
	for _, v := range ints {
		if v == m {
			largest++
		}
	}
	return largest
}
