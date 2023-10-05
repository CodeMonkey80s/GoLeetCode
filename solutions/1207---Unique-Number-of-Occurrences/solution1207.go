package solution1207

import (
	"math"
)

// ============================================================================
// 1207. Unique Number of Occurrences
// URL: https://leetcode.com/problems/unique-number-of-occurrences/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1207---Unique-Number-of-Occurrences
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_unique-24    	 4219881	       246.4 ns/op	      48 B/op       2 allocs/op
	PASS

*/

func uniqueOccurrences(arr []int) bool {
	freq := make(map[int]int, len(arr))
	for _, val := range arr {
		_, ok := freq[val]
		if !ok {
			freq[val] = 1
		} else {
			freq[val]++
		}
	}
	sl := make([]int, 2002)
	for i, v := range freq {
		if v < 0 {
			v = int(math.Abs(float64(v))) + 1
		} else {
			v += 1001
		}
		if sl[v] != 0 {
			return false
		}
		sl[v] = i
	}
	return true
}
