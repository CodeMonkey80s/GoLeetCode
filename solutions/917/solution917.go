package solution917

// ============================================================================
// 917. Reverse Only Letters
// URL: https://leetcode.com/problems/reverse-only-letters/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/917
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_reverseOnlyLetters
	Benchmark_reverseOnlyLetters-24    	93101930	        12.38 ns/op	       5 B/op	       1 allocs/op
	PASS
*/

func reverseOnlyLetters(s string) string {

	isLetter := func(ch byte) bool {
		switch {
		case ch >= 'a' && ch <= 'z':
			return true
		case ch >= 'A' && ch <= 'Z':
			return true
		default:
			return false
		}
	}

	a := 0
	b := len(s) - 1

	st := []byte(s)

	for a < b {
		switch {
		case isLetter(st[a]) && isLetter(st[b]):
			st[a], st[b] = st[b], st[a]
			a++
			b--
		case isLetter(st[a]) && !isLetter(st[b]):
			b--
		case !isLetter(st[a]) && isLetter(st[b]):
			a++
		default:
			a++
			b--
		}
	}

	return string(st)
}
