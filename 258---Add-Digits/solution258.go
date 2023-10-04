package solution258

// ============================================================================
// 258. Add Digits
// URL: https://leetcode.com/problems/add-digits/
// ============================================================================

func addDigits(num int) int {
	sum := 0
	var i int = 0
	for {
		for num > 0 {
			sum = sum + int(num%10)
			num = int(num / 10)
			i++
		}
		num = sum
		if num <= 9 {
			break
		}
		sum = 0
	}
	return sum
}
