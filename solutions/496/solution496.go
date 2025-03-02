package solution496

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/496
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_nextGreaterElementV2
	Benchmark_nextGreaterElementV2-24    	 8935968	       112.7 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func nextGreaterElementV2(nums1 []int, nums2 []int) []int {

	indexes := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		indexes[nums2[i]] = i
	}

	for i := 0; i < len(nums1); i++ {
		num := nums1[i]
		idx := indexes[num]
		highest := -1
		for j := idx; j < len(nums2); j++ {
			if nums2[j] > num {
				highest = nums2[j]
				break
			}
		}
		nums1[i] = highest
	}

	return nums1
}

func nextGreaterElementV1(nums1 []int, nums2 []int) []int {

	indexes := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		indexes[nums2[i]] = i
	}

	nextGreater := make(map[int]int)
	stack := make([]int, len(nums1))
	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			nextGreater[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}

	output := make([]int, len(nums1))
	for i, num := range nums1 {
		if val, found := nextGreater[num]; found {
			output[i] = val
		} else {
			output[i] = -1
		}
	}

	return output
}
