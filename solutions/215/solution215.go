package solution215

import (
	"slices"
)

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/215
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_findKthLargestV1
	Benchmark_findKthLargestV1-24    	100000000	        15.79 ns/op	       0 B/op	       0 allocs/op
	Benchmark_findKthLargestV2
	Benchmark_findKthLargestV2-24    	13862401	        73.57 ns/op	      48 B/op	       3 allocs/op
	PASS
*/

func findKthLargestV2(nums []int, k int) int {

	stack := make([]int, k)
	for i := range k {
		stack[i] = nums[i]
	}
	slices.Sort(stack)

	for i := k; i < len(nums); i++ {

		if nums[i] > stack[0] {
			stack = stack[1:]
			stack = append(stack, nums[i])
			slices.Sort(stack)
		}

	}

	return stack[len(stack)-k]
}

func findKthLargestV1(nums []int, k int) int {
	slices.SortFunc(nums, func(a, b int) int {
		return b - a
	})

	return nums[k-1]
}
