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
	Benchmark_stringSequence-24    	 2483098	       420.8 ns/op	     498 B/op	      17 allocs/op
	PASS
*/

func stringSequence(target string) []string {

	sequence := []string{""}

	var n int
	var prefix string
	for _, ch := range target {
		n = int(ch-'a') + 1
		prefix = sequence[len(sequence)-1]
		for i := 0; i < n; i++ {
			sequence = append(sequence, prefix+string(rune(i+'a')))

		}
	}

	return sequence[1:]
}
