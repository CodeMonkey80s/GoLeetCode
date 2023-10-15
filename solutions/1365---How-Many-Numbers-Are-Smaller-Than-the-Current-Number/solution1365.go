package solution1365

// ============================================================================
// 1365. How Many Numbers Are Smaller Than the Current Number
// URL: https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1365---How-Many-Numbers-Are-Smaller-Than-the-Current-Number
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_smallerNumbers-24    	45587346	        32.01 ns/op	      48 B/op	       1 allocs/op
	PASS

*/

func smallerNumbersThanCurrent(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i <= len(nums)-1; i++ {
		c := 0
		for j := 0; j <= len(nums)-1; j++ {
			if i != j {
				if nums[j] < nums[i] {
					c++
				}
			}
		}
		ans[i] = c
	}
	return ans
}
