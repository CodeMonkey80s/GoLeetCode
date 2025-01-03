package solution2451

// ============================================================================
// 2451. Odd String Difference
// URL: https://leetcode.com/problems/odd-string-difference/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2451---Odd-String-Difference
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_oddString-24    	 3260404	       388.6 ns/op	    1792 B/op	       2 allocs/op
	PASS

*/

func oddString(words []string) string {
	type pair struct {
		s string
		n int
	}
	type hash [101]int
	m := make(map[hash]pair, len(words)-1)
	for _, word := range words {
		values := hash{}
		for i := 0; i < len(word)-1; i++ {
			a := int(word[i] - 'a')
			b := int(word[i+1] - 'a')
			diff := b - a
			values[i] = diff
		}
		_, ok := m[values]
		if !ok {
			m[values] = pair{
				s: word,
				n: 1,
			}
		} else {
			p := m[values]
			p.n++
			m[values] = p
		}
	}
	for _, p := range m {
		if p.n == 1 {
			return p.s
		}
	}
	return ""
}
