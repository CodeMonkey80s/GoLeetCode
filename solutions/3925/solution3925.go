package solution3925

func concatWithReverse(nums []int) []int {
	reversed := make([]int, len(nums)*2)
	for i := 0; i < len(nums); i++ {
		a := i
		b := (2 * len(nums)) - 1 - i
		reversed[a], reversed[b] = nums[i], nums[i]
	}

	return reversed
}
