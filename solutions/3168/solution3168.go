package solution3168

// ============================================================================
// 3168. Minimum Number of Chairs in a Waiting Room
// URL: https://leetcode.com/problems/minimum-number-of-chairs-in-a-waiting-room/description/
// ============================================================================

func minimumChairs(s string) int {
	n := 0
	chairs := 0
	for _, ch := range s {
		switch ch {
		case 'E':
			n++
		case 'L':
			n--
		}
		if n > chairs {
			chairs = n
		}
	}

	return chairs
}
