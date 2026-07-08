package solution3701

func alternatingSum(nums []int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		switch {
		// Even
		case i%2 == 0:
			sum += nums[i]
		// Odd
		default:
			sum -= nums[i]
		}
	}

	return sum
}
