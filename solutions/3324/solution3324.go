package solution3324

// ============================================================================
// 3324. Find the Sequence of Strings Appeared on the Screen
// URL: https://leetcode.com/problems/find-the-sequence-of-strings-appeared-on-the-screen/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3324
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_stringSequence
	Benchmark_stringSequence-24    	 2006986	       564.8 ns/op	     568 B/op	      23 allocs/op
	PASS
*/

func stringSequence(target string) []string {

	var sequence []string
	for _, ch := range target {

		n := int(ch-'a') + 1
		prefix := ""
		if len(sequence) >= 1 {
			l := len(sequence[len(sequence)-1])
			prefix += sequence[len(sequence)-1][:l]
		}

		for i := 0; i < n; i++ {
			s := prefix + ""
			s += string(rune(i + 'a'))
			sequence = append(sequence, s)

		}
	}

	return sequence
}
