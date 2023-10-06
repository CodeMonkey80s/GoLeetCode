package solution415

import (
	"math/big"
	"strings"
)

// ============================================================================
// 415. Add Strings
// URL: https://leetcode.com/problems/add-strings/
// ============================================================================

/*
	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/415---Add-Strings
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_addStrings-24        	16012267	        87.12 ns/op	      28 B/op	       7 allocs/op
	Benchmark_addStrings_std-24    	 5277220	       227.2 ns/op	     136 B/op	       7 allocs/op
	PASS
*/

func addStrings(num1 string, num2 string) string {
	if !(1 <= len(num1) && len(num1) <= 10_000) {
		return "0"
	}
	if !(1 <= len(num2) && len(num2) <= 10_000) {
		return "0"
	}
	if num1 == "0" && num2 == "0" {
		return "0"
	}
	if num2 == "0" {
		return num1
	}
	if num1 == "0" {
		return num2
	}
	ans := ""
	a := len(num1) - 1
	b := len(num2) - 1
	v1 := 0
	v2 := 0
	carry := 0
	digit := ""
	for a >= -1 || b >= -1 {
		v1 = 0
		if a >= 0 {
			v1 = int(num1[a] - '0')
		}
		v2 = 0
		if b >= 0 {
			v2 = int(num2[b] - '0')
		}
		sum := v1 + v2 + carry
		if sum >= 10 {
			carry = 1
			digit = string(rune(sum-10) + '0')
		} else {
			carry = 0
			digit = string(rune(sum) + '0')
		}
		ans = digit + ans
		a--
		b--
	}
	return strings.TrimLeft(ans, "0")
}

func addStrings_std(num1 string, num2 string) string {
	v1 := big.Int{}
	v1.SetString(num1, 10)
	v2 := big.Int{}
	v2.SetString(num2, 10)
	sum := &big.Int{}
	sum = sum.Add(&v1, &v2)

	return sum.String()
}
