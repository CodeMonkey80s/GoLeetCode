package solution2828

// ============================================================================
// 2828. Check if a String is an Acronym of Words
// URL: https://leetcode.com/problems/check-if-a-string-is-an-acronym-of-words/description/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2828---Check-if-a-String-is-an-Acronym-of-Words
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_isAcronym-24    	665971717	         1.805 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func isAcronym(words []string, s string) bool {
	if len(s) < len(words) || len(s) > len(words) {
		return false
	}
	for i := 0; i < len(words); i++ {
		if words[i][0] != s[i] {
			return false
		}
	}
	return true
}
