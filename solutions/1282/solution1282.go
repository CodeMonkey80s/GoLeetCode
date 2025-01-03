package solution1282

// ============================================================================
// 1282. Group the People Given the Group Size They Belong To
// URL: https://leetcode.com/problems/group-the-people-given-the-group-size-they-belong-to/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1282
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_groupThePeople
	Benchmark_groupThePeople-24    	 3862311	       310.3 ns/op	     296 B/op	       8 allocs/op
	PASS
*/

func groupThePeople(groupSizes []int) [][]int {

	freq := make(map[int][]int)
	for id, size := range groupSizes {
		_, ok := freq[size]
		if !ok {
			freq[size] = make([]int, 0)
		}
		freq[size] = append(freq[size], id)
	}

	var output [][]int
	for size, ids := range freq {
		steps := len(ids) / size
		for i := 0; i < steps; i++ {
			a := i * size
			b := (i + 1) * size
			output = append(output, ids[a:b])
		}
	}

	return output
}
