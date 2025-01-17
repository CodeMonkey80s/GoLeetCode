package solution3360

// ============================================================================
// 3360. Stone Removal Game
// URL: https://leetcode.com/problems/stone-removal-game/
// ============================================================================

func canAliceWin(n int) bool {

	s := 10
	a := true

	for n > 0 {
		n -= s
		s--
		a = !a
	}

	if n == 0 {
		return !a
	}

	return a
}
