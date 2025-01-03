package solution3285

// ============================================================================
// 3285. Find Indices of Stable Mountains
// URL: https://leetcode.com/problems/find-indices-of-stable-mountains/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3285--Find-Indices-of-Stable-Mountains
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_StableMountains
	Benchmark_StableMountains-24    	46737021	        22.68 ns/op	      48 B/op	       1 allocs/op
	PASS
*/

func stableMountains(height []int, threshold int) []int {
	indices := make([]int, 0, len(height))
	for i := 1; i < len(height); i++ {
		if height[i-1] > threshold {
			indices = append(indices, i)
		}
	}

	return indices
}
