package solution119

// ============================================================================
// 119. Pascal's Triangle II
// URL: https://leetcode.com/problems/pascals-triangle-ii/
// ============================================================================

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	out := [][]int{}
	a, b, s := 0, 0, 0
	for d := 0; d <= rowIndex; d++ {
		s = d + 1
		inn := make([]int, s)
		for i := 0; i < s; i++ {
			inn[i] = 1
		}
		for i := 0; i < s-2; i++ {
			a = out[d-1][i]
			b = out[d-1][i+1]
			inn[i+1] = a + b
		}
		out = append(out, inn)
	}
	return out[len(out)-1]
}
