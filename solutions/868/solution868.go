package solution868

// ============================================================================
// 868. Binary Gap
// URL: https://leetcode.com/problems/binary-gap/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/868---Binary-Gap
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	BenchmarkBinaryGap
	BenchmarkBinaryGap-24    	 9231778	       129.9 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func binaryGap(n int) int {

	var a, b, m int

	isOne := func(pos int) bool {
		if val := n & (1 << pos); val > 0 {
			return true
		}

		return false
	}

	a = 63
	b = a - 1
	for {
		switch {
		case isOne(a) && isOne(b):
			m = max(m, a-b)
			a = b
			b--
		case isOne(a) && !isOne(b):
			b--
		case !isOne(a) && !isOne(b):
			b--
			a--
		case !isOne(a) && isOne(b):
			a = b
			b--
		}

		if b < 0 {
			break
		}
	}

	return m
}
