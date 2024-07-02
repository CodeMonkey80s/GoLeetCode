package solution3162

// ============================================================================
// 3162. Find the Number of Good Pairs I
// URL: https://leetcode.com/problems/find-the-number-of-good-pairs-i/
// ============================================================================

func numberOfPairs(nums1 []int, nums2 []int, k int) int {
	total := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i]%(nums2[j]*k) == 0 {
				total++
			}
		}
	}

	return total
}
