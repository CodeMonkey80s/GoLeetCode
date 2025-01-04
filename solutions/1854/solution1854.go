package solution1854

// ============================================================================
// 1854. Maximum Population Year
// URL: https://leetcode.com/problems/maximum-population-year/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1854
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_maximumPopulation
	Benchmark_maximumPopulation-24    	17257236	        68.27 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func maximumPopulations(logs [][]int) int {

	const (
		loRange = 1950
		hiRange = 2050
	)

	population := make([]int, 100)
	for _, log := range logs {
		for i := log[0]; i < log[1]; i++ {
			year := i - loRange
			population[year]++
		}
	}

	var pop int
	var year int
	for y, n := range population {
		if n > pop {
			pop = n
			year = y
		}
	}

	return year + loRange
}
