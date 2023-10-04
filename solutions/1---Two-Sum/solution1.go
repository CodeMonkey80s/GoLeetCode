package solution1

// ============================================================================
// 1. Two Sum
// URL: https://leetcode.com/problems/two-sum/
// ============================================================================

func twoSum_Iterations(nums []int, target int) []int {
	max := len(nums)
	if max < 2 || max > 10_000 {
		return []int{}
	}
	if max == 2 && nums[0]+nums[1] == target {
		return []int{0, 1}
	}
	for i := 0; i < max; i++ {
		for j := i + 1; j < max; j++ {
			t := nums[i] + nums[j]
			if t == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// ============================================================================

func twoSum_Map(nums []int, target int) []int {
	max := len(nums)
	if max < 2 || max > 10_000 {
		return []int{}
	}
	if max == 2 && nums[0]+nums[1] == target {
		return []int{0, 1}
	}
	var t int
	m := make(map[int]int, max)
	for k, v := range nums {
		t = target - v
		_, ok := m[t]
		if ok {
			return []int{m[t], k}
		}
		m[v] = k
	}
	return []int{}
}
