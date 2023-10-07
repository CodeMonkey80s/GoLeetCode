package solution2446

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
	Benchmark_haveConflict-24    	255966577	         4.455 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func haveConflict(event1 []string, event2 []string) bool {
	getMinutes := func(s string) int {
		h := int(s[0]-'0')*10 + (int(s[1] - '0'))
		m := int(s[3]-'0')*10 + (int(s[4] - '0'))
		v := h*60 + m
		return v
	}
	a1 := getMinutes(event1[0])
	a2 := getMinutes(event1[1])
	b1 := getMinutes(event2[0])
	b2 := getMinutes(event2[1])
	overlap := func(a1, a2, b1, b2 int) bool {
		return b2 >= a1 && a2 >= b1
	}
	return overlap(a1, a2, b1, b2) || overlap(b1, b2, a1, a2)
}
