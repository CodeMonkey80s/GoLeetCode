package solution1550

// ============================================================================
// 1550. Three Consecutive Odds
// URL: https://leetcode.com/problems/three-consecutive-odds/
// ============================================================================

/*

	$ go test -bench=. -benchmem ./...
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1550---Three-Consecutive-Odds
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_threeConsecutiveOdds-24    	916179416	         1.324 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func threeConsecutiveOdds(arr []int) bool {
	for i := 0; i < len(arr)-2; i++ {
		a := arr[i]
		b := arr[i+1]
		c := arr[i+2]
		if a%2 != 0 && b%2 != 0 && c%2 != 0 {
			return true
		}
	}
	return false
}
