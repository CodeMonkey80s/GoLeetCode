package solution575

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/575
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_distributeCandies
	Benchmark_distributeCandies-24    	26258036	        47.47 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func distributeCandies(candyType []int) int {

	freq := make(map[int]struct{}, len(candyType))
	for i := 0; i < len(candyType); i++ {
		freq[candyType[i]] = struct{}{}
	}

	return min(len(candyType)/2, len(freq))
}

func distributeCandiesV1(candyType []int) int {

	n := len(candyType)
	types := n / 2

	freq := make(map[int]int, len(candyType))
	for i := 0; i < n; i++ {
		freq[candyType[i]]++
	}

	if types < len(freq) {
		return types
	}

	return len(freq)
}
