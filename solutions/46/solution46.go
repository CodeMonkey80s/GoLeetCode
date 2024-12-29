package solution46

// ============================================================================
// 46, Permutations
// URL: https://leetcode.com/problems/permutations/
// ============================================================================

import "slices"

var perms [][]int

func permute(nums []int) [][]int {

	perms = make([][]int, 0)
	heapPermutation(nums, len(nums), len(nums))

	return perms
}

func heapPermutation(a []int, size int, n int) {
	if size == 1 {
		if len(a[:n]) >= 1 {
			perms = append(perms, slices.Clone(a[:n]))
		}
		return
	}

	for i := 0; i < size; i++ {
		heapPermutation(a, size-1, n)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}
