package solution3386

// ============================================================================
// 3386. Button with Longest Push Time
// URL: https://leetcode.com/problems/button-with-longest-push-time/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3386
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_buttonWithLongestTime
	Benchmark_buttonWithLongestTime-24    	100000000	        10.48 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func buttonWithLongestTime(events [][]int) int {

	var t int
	var id int
	var longest int
	for i := 0; i <= len(events)-1; i++ {
		if i == 0 {
			t = events[i][1]
		} else {
			t = events[i][1] - events[i-1][1]
		}

		switch {
		case t > longest:
			longest = t
			if id == 0 || events[i][0] != id {
				id = events[i][0]
			}
		case t == longest:
			if id == 0 || events[i][0] < id {
				id = events[i][0]
			}
		}
	}

	return id
}
