package solution2248

import "slices"

func intersection(nums [][]int) []int {

	var m int
	freq := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m = max(m, len(nums[i]))
		for j := 0; j < len(nums[i]); j++ {
			freq[nums[i][j]]++
		}
	}

	output := make([]int, 0, m)
	for number, count := range freq {
		if count == len(nums) {
			output = append(output, number)
		}
	}

	slices.Sort(output)

	return output
}
