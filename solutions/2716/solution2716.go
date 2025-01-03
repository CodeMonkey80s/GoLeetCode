package solution2716

// ============================================================================
// 2716. Minimize String Length
// URL: https://leetcode.com/problems/minimize-string-length/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2716---Minimize-String-Length
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_minimizeStringLength-24    	362493421	         3.318 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func minimizedStringLength(s string) int {
	freq := make([]int, 26)
	ans := 0
	idx := 0
	for _, char := range s {
		idx = int(char - 'a')
		if freq[idx] == 0 {
			ans++
		}
		freq[idx]++
	}
	return ans
}
