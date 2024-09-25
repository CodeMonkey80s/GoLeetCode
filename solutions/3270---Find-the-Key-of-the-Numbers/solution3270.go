package solution3270

import (
	"fmt"
	"strconv"
	"strings"
)

// ============================================================================
// 3270. Find the Key of the Numbers
// URL: https://leetcode.com/problems/find-the-key-of-the-numbers/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3270---Find-the-Key-of-the-Numbers
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	BenchmarkGenerateKeyV2
	BenchmarkGenerateKeyV2-24    	84724898	        14.52 ns/op	       0 B/op	       0 allocs/op
	BenchmarkGenerateKeyV1
	BenchmarkGenerateKeyV1-24    	 4832989	       222.3 ns/op	      32 B/op	       5 allocs/op
	PASS
*/

func generateKeyV2(num1 int, num2 int, num3 int) int {

	ans := []int{0, 0, 0, 0}

	getDigits := func(num int) []int {
		digits := make([]int, 4)
		i := 3
		for num > 0 {
			digits[i] = num % 10
			num /= 10
			i--
		}
		return digits
	}

	d1 := getDigits(num1)
	d2 := getDigits(num2)
	d3 := getDigits(num3)

	for i := 0; i < 4; i++ {
		ans[i] = min(d1[i], d2[i], d3[i])
	}

	return ans[0]*1000 + ans[1]*100 + ans[2]*10 + ans[3]
}

func generateKeyV1(num1 int, num2 int, num3 int) int {

	var sb strings.Builder

	n1 := fmt.Sprintf("%04d", num1)
	n2 := fmt.Sprintf("%04d", num2)
	n3 := fmt.Sprintf("%04d", num3)

	for i := 0; i < 4; i++ {
		d := min(n1[i], n2[i], n3[i])
		sb.WriteString(strconv.Itoa(int(d)))
	}

	val, _ := strconv.Atoi(sb.String())

	return val
}
