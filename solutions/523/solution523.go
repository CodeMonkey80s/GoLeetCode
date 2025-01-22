package solution523

func checkSubarraySum(nums []int, k int) bool {

	remainderMap := make(map[int]int)
	remainderMap[0] = -1
	cumulativeSum := 0

	for i, num := range nums {
		cumulativeSum += num
		remainder := cumulativeSum % k

		if prevIndex, exists := remainderMap[remainder]; exists {
			if i-prevIndex > 1 {
				return true
			}
		} else {
			remainderMap[remainder] = i
		}
	}

	return false
}
