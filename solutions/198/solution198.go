package solution198

// ============================================================================
// 198. House Robber
// https://leetcode.com/problems/house-robber/
// ============================================================================

func rob(nums []int) int {
	rob1, rob2, t := 0, 0, 0
	for _, v := range nums {
		if v+rob1 > rob2 {
			t = v + rob1
		} else {
			t = rob2
		}
		rob1 = rob2
		rob2 = t
	}
	return rob2
}
