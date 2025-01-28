package solution1725

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1725
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countGoodRectangles
	Benchmark_countGoodRectangles-24    	500781232	         5.642 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func countGoodRectangles(rectangles [][]int) int {

	var highest int
	var count int
	var v int
	for i := 0; i < len(rectangles); i++ {
		v = min(rectangles[i][0], rectangles[i][1])
		switch {
		case v == highest:
			count++
		case v > highest:
			highest = v
			count = 1
		}
	}

	return count
}
