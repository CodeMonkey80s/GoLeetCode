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
	BenchmarkGenerateKey
	BenchmarkGenerateKey-24    	 4511347	       252.6 ns/op	      32 B/op	       5 allocs/op
	PASS
*/

func generateKey(num1 int, num2 int, num3 int) int {

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
