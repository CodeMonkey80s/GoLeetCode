package solution1389

import "slices"

// ============================================================================
// 1389. Create Target Array in the Given Order
// URL: https://leetcode.com/problems/create-target-array-in-the-given-order/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1389---Create-Target-Array-in-the-Given-Order
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_createTargetArray-24        	68775423	        17.50 ns/op	       0 B/op	       0 allocs/op
	Benchmark_createTargetArray_std-24    	11325408	       109.1 ns/op	     896 B/op	       1 allocs/op
	PASS

*/

func createTargetArray(nums []int, index []int) []int {
	ans := make([]int, 0, 100)
	insert := func(a []int, index int, value int) []int {
		if len(a) == index {
			return append(a, value)
		}
		a = append(a[:index+1], a[index:]...)
		a[index] = value
		return a
	}
	for i, v := range nums {
		ans = insert(ans, index[i], v)
	}
	return ans
}

// At the time of writing the LeetCode does not support the "slices" package...
func createTargetArray_std(nums []int, index []int) []int {
	ans := make([]int, 0, 100)
	for i, v := range nums {
		ans = slices.Insert(ans, index[i], v)
	}
	return ans
}
