package solution2446

import (
	"strconv"
)

// ============================================================================
// 2446. Determine if Two Events Have Conflict
// URL: https://leetcode.com/problems/determine-if-two-events-have-conflict/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2446---Determine-if-Two-Events-Have-Conflict
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_haveConflict-24    	58349853	        20.16 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func haveConflict(event1 []string, event2 []string) bool {
	getMinutes := func(s string) int {
		h, err := strconv.Atoi(s[:2])
		if err != nil {
			panic(err)
		}
		m, err := strconv.Atoi(s[3:])
		if err != nil {
			panic(err)
		}
		v := h*60 + m
		return v
	}
	eventOneTimeA := getMinutes(event1[0])
	eventOneTimeB := getMinutes(event1[1])
	eventTwoTimeA := getMinutes(event2[0])
	eventTwoTimeB := getMinutes(event2[1])
	if eventTwoTimeA <= eventOneTimeB && eventTwoTimeB >= eventOneTimeA {
		return true
	}

	if eventTwoTimeB >= eventOneTimeA && eventTwoTimeB <= eventOneTimeB {
		return true
	}

	if eventTwoTimeA >= eventOneTimeA && eventTwoTimeB <= eventOneTimeB {
		return true
	}
	if eventOneTimeA >= eventTwoTimeA && eventOneTimeB <= eventTwoTimeB {
		return true
	}
	return false
}
