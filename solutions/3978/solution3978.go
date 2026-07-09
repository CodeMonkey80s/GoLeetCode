package solution3978

func isMiddleElementUnique(nums []int) bool {
	middle := nums[len(nums)/2]
	for i := 0; i < len(nums); i++ {
		if i != len(nums)/2 && nums[i] == middle {
			return false
		}
	}

	return true
}
