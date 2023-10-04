package solution1704

// ============================================================================
// 1704. Determine if String Halves Are Alike
// URL: https://leetcode.com/problems/determine-if-string-halves-are-alike/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_halvesAreAlike-24    	 4612999	       259.8 ns/op	     179 B/op	       1 allocs/op
	PASS

*/

func halvesAreAlike(s string) bool {
	size := len(s)
	s1 := s[0 : size/2]
	s2 := s[size/2:]
	c1 := 0
	c2 := 0
	m := map[byte]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
		'A': 1,
		'E': 1,
		'I': 1,
		'O': 1,
		'U': 1,
	}
	for _, ch := range s1 {
		_, ok := m[byte(ch)]
		if ok {
			c1++
		}
	}
	for _, ch := range s2 {
		_, ok := m[byte(ch)]
		if ok {
			c2++
		}
	}
	return c1 == c2
}
