package solution2119

import (
	"strconv"
	"strings"
)

// ============================================================================
// 2119. A Number After a Double Reversal
// URL: https://leetcode.com/problems/a-number-after-a-double-reversal/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2119---A-Number-After-a-Double-Reversal
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	BenchmarkIsSameAfterReversals
	BenchmarkIsSameAfterReversals-24    	 5862866	       177.8 ns/op	      56 B/op	       5 allocs/op
	PASS
*/

func isSameAfterReversals(num int) bool {
	if num == 0 {
		return true
	}

	reverse := func(rns []rune) []rune {
		var sb strings.Builder
		for i := len(rns) - 1; i >= 0; i-- {
			sb.WriteRune(rns[i])
		}

		return []rune(sb.String())
	}

	s := []rune(strconv.Itoa(num))
	r1 := reverse(s)
	r1 = []rune(strings.TrimLeft(string(r1), "0"))
	r2 := reverse(r1)

	return string(r2) == string(s)
}
