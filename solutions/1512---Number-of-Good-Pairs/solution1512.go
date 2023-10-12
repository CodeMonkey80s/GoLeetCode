package solution1512

// ============================================================================
// 1512. Number of Good Pairs
// URL: https://leetcode.com/problems/number-of-good-pairs/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1512---Number-of-Good-Pairs
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_squareIsWhite-24    	228702776	         4.615 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func numIdenticalPairs(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				ans++
			}
		}
	}
	return ans
}
