package solution442

// ============================================================================
// 442. Find All Duplicates in an Array
// URL: https://leetcode.com/problems/find-all-duplicates-in-an-array/
// ============================================================================

func findDuplicates(nums []int) []int {
	freq := make(map[int]int, len(nums))
	duplicates := make([]int, 0, len(nums))
	for _, num := range nums {
		freq[num]++
		if freq[num] == 2 {
			duplicates = append(duplicates, num)
		}
	}

	return duplicates
}
