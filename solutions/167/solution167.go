package solution167

// ============================================================================
// 167. Two Sum II
// URL: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_twoSum-24    	97250382	        14.11 ns/op	      16 B/op	       1 allocs/op
	PASS

*/

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	sum := 0
	for l < r {
		sum = numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum <= target {
			l++
		} else {
			r--
		}
	}
	return []int{0}
}
