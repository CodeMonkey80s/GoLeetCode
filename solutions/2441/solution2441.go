package solution2441

func findMax(nums []int) int {

	largest := -1
	sums := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {

		value := nums[i]
		switch {
		case value >= 0:
			sums[value] += value
			op := -value
			_, ok := sums[op]
			if ok {
				largest = max(largest, value)
			}
		default:
			sums[value] += value
			op := -value
			_, ok := sums[op]
			if ok {
				largest = max(largest, op)
			}
		}

	}

	return largest
}
