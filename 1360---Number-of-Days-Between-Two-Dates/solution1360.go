package solution1360

import (
	"math"
	"time"
)

// ============================================================================
// 1360. Number of Days Between Two Dates
// URL: https://leetcode.com/problems/number-of-days-between-two-dates/
// ============================================================================

func daysBetweenDates(date1 string, date2 string) int {
	t1, err := time.Parse("2006-01-02", date1)
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse("2006-01-02", date2)
	if err != nil {
		panic(err)
	}
	return int(math.Abs(t1.Sub(t2).Hours() / 24))
}
