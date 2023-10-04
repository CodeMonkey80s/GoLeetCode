package solution88

// ============================================================================
// 88. Merge Sorted Array
// URL: https://leetcode.com/problems/merge-sorted-array/
// ============================================================================

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	for i = 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	for i = 1; i < m+n; i++ {
		for j = 0; j < m+n-1; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
			}
		}
	}
}
