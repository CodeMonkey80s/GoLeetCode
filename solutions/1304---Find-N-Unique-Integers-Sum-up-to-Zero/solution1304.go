package solution1304

// ============================================================================
// i304. Find N Unique Integers Sum up to Zero
// URL: https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1304---Find-N-Unique-Integers-Sum-up-to-Zero
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_sumZero
	Benchmark_sumZero-24    	12513075	        81.94 ns/op	     120 B/op	       4 allocs/op
	PASS
*/

func sumZero(n int) []int {
	var output []int
	for i := 1; i <= n/2; i++ {
		output = append(output, i)
		output = append(output, -i)
	}

	if n%2 != 0 {
		output = append(output, 0)
	}

	return output
}
