package solution1742

// ============================================================================
// 1742. Maximum Number of Balls in a Box
// URL: https://leetcode.com/problems/maximum-number-of-balls-in-a-box/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1742---Maximum-Number-of-Balls-in-a-Box
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countBalls-24    	 3749175	       326.7 ns/op	     336 B/op	       3 allocs/op
	PASS

*/

func countBalls(lowLimit int, highLimit int) int {
	ans := 0
	freq := make(map[int]int)
	sum := func(num int) int {
		res := 0
		for num > 0 {
			res += num % 10
			num /= 10
		}
		return res
	}
	for i := lowLimit; i <= highLimit; i++ {
		n := sum(i)
		freq[n]++
	}
	for _, v := range freq {
		if v > ans {
			ans = v
		}
	}
	return ans
}
