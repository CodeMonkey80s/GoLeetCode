package solution2917

import (
	"strconv"
	"strings"
)

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2917---Find-the-K-or-of-an-Array
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_findKOrV2
	Benchmark_findKOrV2-24    	10514353	       116.1 ns/op	       0 B/op	       0 allocs/op
	Benchmark_findKOrV1
	Benchmark_findKOrV1-24    	  819944	      1405 ns/op	    1616 B/op	       3 allocs/op
	PASS
*/

func findKOrV2(nums []int, k int) int {
	var c, ans int
	for i := 0; i < 32; i++ {
		c = 0
		for _, num := range nums {
			val := num & (1 << i)
			if val > 0 {
				c++
			}
		}
		if c >= k {
			ans |= 1 << i
		}
	}

	return ans
}

func findKOrV1(nums []int, k int) int {

	bits := make([][32]int, 0, len(nums))

	for _, num := range nums {
		var acc [32]int
		for i := 0; i < 32; i++ {
			val := num & (1 << i)
			if val > 0 {
				acc[31-i] = 1
			}
		}
		bits = append(bits, acc)
	}

	var res [32]rune

	c := 0
	for i := 0; i < 32; i++ {
		res[i] = '0'
		c = 0
		for _, a := range bits {
			if a[i] == 1 {
				c++
			}
		}
		if c >= k {
			res[i] = '1'
		}
	}

	var sb strings.Builder
	sb.WriteString(string(res[:]))
	s := sb.String()

	n, _ := strconv.ParseInt(s, 2, 64)

	return int(n)
}
