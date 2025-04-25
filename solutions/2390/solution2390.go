package solution2390

// ============================================================================
// 2390. Removing Stars From a String
// URL: https://leetcode.com/problems/removing-stars-from-a-string/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2390
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_removeStars
	Benchmark_removeStars-24    	65154109	        26.72 ns/op	      16 B/op	       1 allocs/op
	PASS
*/

func removeStars(s string) string {

	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '*':
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
