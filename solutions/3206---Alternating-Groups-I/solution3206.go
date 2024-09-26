package solution3206

// ============================================================================
// 3206. Alternating Groups I
// URL: https://leetcode.com/problems/alternating-groups-i/
// ============================================================================

func numberOfAlternatingGroups(colors []int) int {
	var a, b, c, n int

	for i := 0; i < len(colors); i++ {
		a = i
		b = i + 1
		c = i + 2

		if b >= len(colors) {
			b = i - len(colors) + 1
		}
		if c >= len(colors) {
			c = i - len(colors) + 2
		}

		if colors[a] == 1 && colors[b] == 0 && colors[c] == 1 || colors[a] == 0 && colors[b] == 1 && colors[c] == 0 {
			n++
		}
	}

	return n
}
