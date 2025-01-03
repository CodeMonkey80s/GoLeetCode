package solution2744

// ============================================================================
// 2744. Find Maximum Number of String Pairs
// URL: https://leetcode.com/problems/find-maximum-number-of-string-pairs/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2744---Find-Maximum-Number-of-String-Pairs
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_maximumNumberOfStringPairs-24                  	33142444	        45.46 ns/op	      16 B/op	       1 allocs/op
	Benchmark_maximumNumberOfStringPairs_first_attempt-24    	 3972052	       343.7 ns/op	      34 B/op	      17 allocs/op
	PASS

*/

func maximumNumberOfStringPairs(words []string) int {
	ans := 0
	type pair [2]byte
	sl := make([]pair, 0, len(words))
	for _, word := range words {
		p := pair{0: word[0], 1: word[1]}
		sl = append(sl, p)
	}
outer:
	for i, word1 := range sl {
		for j, word2 := range sl {
			if j > i {
				p := pair{0: word2[1], 1: word2[0]}
				if word1 == p {
					sl = append(sl[:i], sl[i:]...)
					ans++
					continue outer
				}
			}
		}
	}
	return ans
}

func maximumNumberOfStringPairs_first_attempt(words []string) int {
	ans := 0
	pairs := make(map[string]string, len(words))
	rev := func(s []byte) string {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return string(s)
	}
outer:
	for i := 0; i < len(words); i++ {
		word1 := words[i]
		_, ok1 := pairs[word1]
		if !ok1 {
			pairs[word1] = ""
		}
		for j := 1; j < len(words); j++ {
			word2 := rev([]byte(words[j]))
			if word1 == word2 && j > i {
				pairs[word1] = word2
				ans++
				continue outer
			}
		}
	}
	return ans
}
