package solution2149

// ============================================================================
// 2149. Rearrange Array Elements by Sign
// URL: https://leetcode.com/problems/rearrange-array-elements-by-sign/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2149
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_rearrangeArray
	Benchmark_rearrangeArray-24    	16609581	        65.90 ns/op	      96 B/op	       3 allocs/op
	PASS
*/

func rearrangeArray(nums []int) []int {

	pos := make([]int, 0, len(nums)/2)
	neg := make([]int, 0, len(nums)/2)
	for _, n := range nums {
		switch {
		case n < 0:
			neg = append(neg, n)
		default:
			pos = append(pos, n)
		}
	}

	var a int
	var b int
	out := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		switch i % 2 {
		case 0:
			out = append(out, pos[a])
			a++
		default:
			out = append(out, neg[b])
			b++
		}
	}

	return out

}
