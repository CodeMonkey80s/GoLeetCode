package solution171

import (
	"math"
)

// ============================================================================
// 168. Excel Sheet Column Number
// URL: https://leetcode.com/problems/excel-sheet-column-number/
//
// HINT: https://www.geeksforgeeks.org/program-for-hexadecimal-to-decimal/
// ============================================================================

func titleToNumber(columnTitle string) int {
	var num int
	var cv int
	var ch byte
	var pow float64
	size := len(columnTitle)
	for i := 0; i < size; i++ {
		ch = columnTitle[i]
		pow = float64(size - i - 1)
		if ch >= 'A' && ch <= 'Y' {
			cv = int(ch) - 64
		} else {
			cv = 26
		}
		num += int(float64(cv) * math.Pow(26, pow))
	}
	return num
}
