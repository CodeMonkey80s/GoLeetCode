package solution168

// ============================================================================
// 168. Excel Sheet Column Title
// URL: https://leetcode.com/problems/excel-sheet-column-title/
// ============================================================================

func convertToTitle(columnNumber int) string {
	var s []byte
	r := 0
	n := columnNumber
	for n > 0 {
		r = (n - 1) % 26
		s = append([]byte{byte(65 + r)}, s...)
		n = (n - 1) / 26
	}

	return string(s)
}
