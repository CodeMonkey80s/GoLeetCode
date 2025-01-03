package solution551

import "strings"

// ============================================================================
// 551. Student Attendance Record I
// URL: https://leetcode.com/problems/student-attendance-record-i/
// ============================================================================

func checkRecord(s string) bool {
	return !strings.Contains(s, "LLL") && strings.Count(s, "A") < 2
}
