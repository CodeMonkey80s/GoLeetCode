package solution2810

// ============================================================================
// 2810. Faulty Keyboard
// URL: https://leetcode.com/problems/faulty-keyboard/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2810---Faulty-Keyboard
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_finalString-24    	37700104	        31.27 ns/op	      16 B/op	       2 allocs/op

*/

func finalString(s string) string {
	buf := make([]byte, 0)
	rev := func(s []byte) []byte {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return s
	}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == 'i' {
			buf = rev(buf)
		} else {
			buf = append(buf, ch)
		}
	}
	return string(buf)
}
