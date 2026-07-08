package solution3688

func evenNumberBitwiseORs(nums []int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			if sum == 0 {
				sum = nums[i]
				continue
			}
			sum |= nums[i]
		}
	}
	return sum
}
