package solution2475

func unequalTripets(nums []int) int {

	var n int
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] == nums[j] {
				continue
			}
			for k := j; k < len(nums); k++ {
				if nums[j] == nums[k] {
					continue
				}
				if nums[i] == nums[k] {
					continue
				}
				n++
			}
		}
	}

	return n
}
