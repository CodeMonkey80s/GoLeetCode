package solution3427

func subarraySum(nums []int) int {

	prefixSum := make([]int, len(nums))

	var sum int
	var total int
	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := max(0, i-nums[i]); j <= i; j++ {
			sum += nums[j]
		}
		prefixSum[i] = sum
		total += prefixSum[i]
	}

	return total
}
