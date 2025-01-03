package solution118

// ============================================================================
// 118. Pascal's Triangle
// URL: https://leetcode.com/problems/pascals-triangle/
// ============================================================================

func generate(numRows int) [][]int {
	out := [][]int{}
	a, b, s := 0, 0, 0
	for d := 0; d < numRows; d++ {
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
	return out
}
