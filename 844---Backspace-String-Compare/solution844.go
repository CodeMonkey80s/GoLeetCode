package solution844

// ============================================================================
// 844. Backspace String Compare
// URL: https://leetcode.com/problems/backspace-string-compare/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_backspaceCompare-24    	29367235	        42.69 ns/op	       4 B/op	       2 allocs/op
	PASS

*/

func backspaceCompare(s string, t string) bool {
	i := 0
	m := len(s)
	for {
		switch {
		case s[i] == '#' && i == 0:
			s = s[0:i] + s[i+1:]
			m = len(s)
			i--
		case s[i] == '#' && i >= 1:
			s = s[0:i-1] + s[i+1:]
			m = len(s)
			i--
		case s[i] == '#' && i == 1:
			s = s[:i]
			m = len(s)
			i--
		default:
			i++
		}
		if i < 0 {
			i = 0
		}
		if i >= m {
			break
		}
	}
	i = 0
	m = len(t)
	for {
		switch {
		case t[i] == '#' && i == 0:
			t = t[0:i] + t[i+1:]
			m = len(t)
			i--
		case t[i] == '#' && i >= 1:
			t = t[0:i-1] + t[i+1:]
			m = len(t)
			i--
		case t[i] == '#' && i == 1:
			t = t[:i]
			m = len(t)
			i--
		default:
			i++
		}
		if i < 0 {
			i = 0
		}
		if i >= m {
			break
		}
	}
	return s == t
}
