package solution2848

// ============================================================================
// 2848. Points That Intersect With Cars
// URL: https://leetcode.com/problems/points-that-intersect-with-cars/
// ============================================================================

/*
	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2848---Points-That-Intersect-With-Cars
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_numberOfPoints-24    	100000000	        14.07 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func numberOfPoints(nums [][]int) int {
	ans := 0
	freq := [100]int{}
	for _, pair := range nums {
		a := pair[0]
		b := pair[1]
		for j := a; j <= b; j++ {
			if freq[j] == 0 {
				ans++
			}
			freq[j] = 1
		}
	}
	return ans
}
