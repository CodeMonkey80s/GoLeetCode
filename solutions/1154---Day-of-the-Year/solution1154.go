package solution1154

import (
	"strconv"
	"time"
)

// ============================================================================
// 1154. Day of the year
// URL: https://leetcode.com/problems/day-of-the-year/
// ============================================================================

func dayOfYear(date string) int {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}

	s := t.Format("002")
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
