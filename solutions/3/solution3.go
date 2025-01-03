package solution3

// ============================================================================
// 3. Longest Substring Without Repeating Characters
// URL: https://leetcode.com/problems/longest-substring-without-repeating-characters/submissions/
// ============================================================================
/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3---Longest-Substring-Without-Repeating-Characters
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_lengthOfLongestSubstring
	Benchmark_lengthOfLongestSubstring-24    	61217971	        27.17 ns/op	       8 B/op	       1 allocs/op
	PASS
*/

func lengthOfLongestSubstring(s string) int {
	sl := make([]byte, 0, len(s))
	var longest int
	for i := 0; i < len(s); i++ {
		sl = append(sl, s[i])
		if i > 0 {
			for j := 0; j < len(sl)-1; j++ {
				if sl[j] == s[i] {
					sl = sl[j+1:]
				}
			}
		}
		if len(sl) > longest {
			longest = len(sl)
		}
	}

	return longest
}
