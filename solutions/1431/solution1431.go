package solution1431

// ============================================================================
// 1431. Kids With the Greatest Number of Candies
// URL: https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1431---Kids-with-the-Greatest-Number-of-Candies
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_kidsWithCandies-24    	100000000	        12.14 ns/op	       5 B/op	       1 allocs/op
	PASS

*/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	ans := make([]bool, len(candies))
	maxval := 0
	for _, v := range candies {
		if v > maxval {
			maxval = v
		}
	}
	for i, v := range candies {
		if v+extraCandies >= maxval {
			ans[i] = true
		}
	}
	return ans
}
