package solution2433

// ============================================================================
// 2433. Find The Original Array of Prefix Xor
// URL: https://leetcode.com/problems/find-the-original-array-of-prefix-xor/
// ============================================================================

func findArray(pref []int) []int {

	output := make([]int, len(pref))
	output[0] = pref[0]
	for i := 1; i <= len(pref)-1; i++ {
		output[i] = pref[i-1] ^ pref[i]

	}

	return output
}
