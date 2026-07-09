package solution3959

import "strconv"

func checkGoodInteger(n int) bool {
	digits := strconv.Itoa(n)

	var digitSum int
	var squareSum int
	for i := 0; i < len(digits); i++ {
		v := int(digits[i] - '0')
		digitSum += v
		squareSum += v * v
	}

	return squareSum-digitSum >= 50
}
