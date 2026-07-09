package solution3931

import "math"

func isAdjacentDiffAtMostTwo(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		a := int(s[i] - '0')
		b := int(s[i+1] - '0')
		if math.Abs(float64(a-b)) > 2 {
			return false
		}
	}

	return true
}
