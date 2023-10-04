package solution541

// ============================================================================
// 541. Reverse String II
// URL: https://leetcode.com/problems/reverse-string-ii/
// ============================================================================

func reverseStr(s string, k int) string {
	sl := []byte(s)
	parts := 2 * k
	for i := 0; i < parts; i += 2 {
		a := i * k
		b := (i + 1) * k
		if b > len(sl) {
			break
		}
		// fmt.Printf("a: %d, b: %d\n", a, b)
		for j := 0; j < k-1; j++ {
			idx1 := a + j
			idx2 := b - j - 1
			// fmt.Printf("%c (%d) <-> %c (%d)\n", sl[idx1], idx1, sl[idx2], idx2)
			sl[idx1], sl[idx2] = sl[idx2], sl[idx1]
		}
		// fmt.Printf("Part: %d:%d\n", a, b)
	}
	// fmt.Printf("SL: %q\n", sl)
	return string(sl)
}
