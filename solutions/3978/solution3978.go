package solution3978

func isMiddleElementUnique(nums []int) bool {
	idx := len(nums) / 2
	val := nums[idx]
	a := idx - 1
	b := idx + 1
	for {
		if a < 0 || b > len(nums) {
			break
		}

		if nums[a] == val || nums[b] == val {
			return false
		}

		a--
		b++
	}

	return true
}
