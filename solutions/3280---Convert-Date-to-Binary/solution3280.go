package solution3280

import (
	"fmt"
	"strconv"
)

// ============================================================================
// 3280. Convert Date to Binary
// URL: https://leetcode.com/problems/convert-date-to-binary/
// ============================================================================

func convertDateToBinary(date string) string {
	a, _ := strconv.Atoi(date[:4])
	b, _ := strconv.Atoi(date[5:7])
	c, _ := strconv.Atoi(date[8:10])
	return fmt.Sprintf("%b-%b-%b", a, b, c)
}
