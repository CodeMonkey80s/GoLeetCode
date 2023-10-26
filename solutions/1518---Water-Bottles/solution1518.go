package solution1518

// ============================================================================
// 1518. Water Bottles
// URL: https://leetcode.com/problems/water-bottles/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1518---Water-Bottles
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_numWaterBottles-24    	41135613	        29.09 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func numWaterBottles(numBottles int, numExchange int) int {
	drank := 0
	empty := 0
	for {
		numBottles--
		empty++
		drank++
		if empty%numExchange == 0 {
			empty = 0
			numBottles++
		}
		if numBottles == 0 {
			break
		}
	}
	return drank
}
