package solution1920

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		idx := nums[i]
		ans[i] = nums[idx]
	}
	return ans
}
