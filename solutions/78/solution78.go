package solution78

func subsets(nums []int) [][]int {

	var sets [][]int

	for i := 0; i < 1<<len(nums); i++ {

		sub := []int{}
		for j := 0; j < len(nums); j++ {
			val := i & (1 << j)
			if val > 0 {
				sub = append(sub, nums[j])
			}
		}

		sets = append(sets, sub)
	}

	return sets
}
