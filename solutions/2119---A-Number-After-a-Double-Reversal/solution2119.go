package solution2119

import (
	"slices"
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
	BenchmarkIsSameAfterReversals-24    	10276513	       117.0 ns/op	      19 B/op	       2 allocs/op
	PASS
*/

func isSameAfterReversals(num int) bool {
	if num == 0 {
		return true
	}

	s := []rune(strconv.Itoa(num))

	r1 := slices.Clone(s)
	slices.Reverse(r1)

	r2 := []rune(strings.TrimLeft(string(r1), "0"))
	slices.Reverse(r2)

	return string(r2) == string(s)
}
