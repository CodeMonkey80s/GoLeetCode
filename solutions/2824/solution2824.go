package solution2824

// ============================================================================
// 2824. Count Pairs Whose Sum is Less than Target
// URL: https://leetcode.com/problems/count-pairs-whose-sum-is-less-than-target/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2824---Count-Pairs-Whose-Sum-is-Less-than-Target
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countPairs-24    	308523226	         3.978 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func countPairs(nums []int, target int) int {
	ans := 0
	for i := 0; i <= len(nums)-1; i++ {
		a := nums[i]
		for j := 1; j <= len(nums)-1; j++ {
			b := nums[j]
			if 0 <= i && i < j && a+b < target {
				ans++
			}
		}
	}
	return ans
}
