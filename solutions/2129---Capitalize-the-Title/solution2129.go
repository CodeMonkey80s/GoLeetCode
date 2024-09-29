package solution2129

// ============================================================================
// 2129. Capitalize the Title
// URL: https://leetcode.com/problems/apple-redistribution-into-boxes/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2129---Capitalize-the-Title
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_capitalizeTitle
	Benchmark_capitalizeTitle-24    	 6810496	       162.1 ns/op	      24 B/op	       1 allocs/op
	PASS
*/

func capitalizeTitle(title string) string {

	s := []rune(title)

	l := 0
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		switch {
		case ch >= 'a' && ch <= 'z':
			l++
		case ch >= 'A' && ch <= 'Z':
			l++
			s[i] = ch + 32
		}

		if ch == ' ' && l > 2 {
			s[i+1] = s[i+1] - 32
			l = 0
		}

		if i == 0 && l > 2 {
			s[i] = s[i] - 32
			l = 0
		}

		if ch == ' ' {
			l = 0
		}

	}

	return string(s)
}
