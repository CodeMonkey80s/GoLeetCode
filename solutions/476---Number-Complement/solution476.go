package solution476

import (
	"math/bits"
	"strconv"
)

// ============================================================================
// 476. Number Complement
// URL: https://leetcode.com/problems/number-complement/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/476---Number-Complement
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_findComplementV3
	Benchmark_findComplementV3-24    	627674793	         1.910 ns/op	   0 B/op	       0 allocs/op
	Benchmark_findComplementV2
	Benchmark_findComplementV2-24    	522780742	         5.248 ns/op	   0 B/op	       0 allocs/op
	Benchmark_findComplementV1
	Benchmark_findComplementV1-24    	16683871	        71.61 ns/op	       8 B/op	       3 allocs/op
	PASS
*/

func findComplementV3(num int) int {
	ans := 0
	for i := 0; num > 0; i++ {
		if num&1 == 0 {
			ans += 1 << i
		}

		num = num >> 1
	}

	return ans
}

func findComplementV2(num int) int {
	ans := num
	for i := 0; i < bits.UintSize-bits.LeadingZeros(uint(num)); i++ {
		val := ans & (1 << i)
		if val == 0 {
			ans |= 1 << i
		} else {
			ans &^= 1 << i
		}
	}
	return ans
}

func findComplementV1(num int) int {
	n1 := strconv.FormatInt(int64(num), 2)
	n2 := ""
	for i := len(n1) - 1; i >= 0; i-- {
		if n1[i] == '0' {
			n2 = "1" + n2
		} else {
			n2 = "0" + n2
		}
	}
	ans, err := strconv.ParseInt(n2, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(ans)
}
