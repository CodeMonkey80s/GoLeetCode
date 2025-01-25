package solution2341

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2341
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_numberOfPairs
	Benchmark_numberOfPairs-24    	 8128572	       167.9 ns/op	      48 B/op	       2 allocs/op
	PASS
*/

func numberOfPairs(nums []int) []int {

	freq := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	var a, b int
	for i := range freq {
		if freq[i] >= 2 {
			a += freq[i] / 2
		}
		if freq[i]%2 != 0 {
			b++
		}
	}

	return []int{a, b}
}
