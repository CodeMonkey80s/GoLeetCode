package solution1920

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1920---Build-Array-from-Permutation
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_buildArray-24    	81729248	        16.63 ns/op	      48 B/op	       1 allocs/op
	PASS

*/

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		idx := nums[i]
		ans[i] = nums[idx]
	}
	return ans
}
