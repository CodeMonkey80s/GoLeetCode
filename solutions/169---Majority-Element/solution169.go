package solution169

// ============================================================================
// 169. Majority Element
// URL: https://leetcode.com/problems/majority-element/
// ============================================================================

func majorityElement(nums []int) int {
	size := len(nums)
	m := make(map[int]int)
	for i := 0; i < size; i++ {
		v := nums[i]
		_, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}
	for k, v := range m {
		if v > size/2 {
			return k
		}
	}
	return 0
}
