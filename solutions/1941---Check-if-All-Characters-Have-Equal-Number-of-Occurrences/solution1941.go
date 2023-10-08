package solution1941

// ============================================================================
// 1941. Check if All Characters Have Equal Number of Occurrences
// URL: https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1941---Check-if-All-Characters-Have-Equal-Number-of-Occurrences
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_areOccurrencesEqual-24    	 3856162	       317.5 ns/op	     752 B/op	       3 allocs/op
	PASS

*/

func areOccurrencesEqual(s string) bool {
	m := make(map[byte]int, 26)
	for _, char := range s {
		ch := byte(char)
		_, ok := m[ch]
		if !ok {
			m[ch] = 1
		} else {
			m[ch]++
		}
	}
	c := -1
	i := 0
	for _, v := range m {
		if i == 0 {
			c = v
		}
		if c != v {
			return false
		}
		i++
	}
	return true
}
