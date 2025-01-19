package solution2442

import (
	"slices"
	"strconv"
	"strings"
)

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2442
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countDistinctIntegersV1
	Benchmark_countDistinctIntegersV1-24    	  732272	      1457 ns/op	    1660 B/op	      26 allocs/op
	Benchmark_countDistinctIntegersV2
	Benchmark_countDistinctIntegersV2-24    	 1784204	       614.3 ns/op	     672 B/op	       2 allocs/op
	PASS
*/

func countDistinctIntegersV2(nums []int) int {

	freq := make(map[int]int, len(nums)*2)

	rev := func(num int) int {
		var res int
		for num > 0 {
			remainder := num % 10
			res = (res * 10) + remainder
			num /= 10
		}
		return res
	}

	for _, num := range nums {
		freq[num]++
		freq[rev(num)]++
	}

	return len(freq)
}

func countDistinctIntegersV1(nums []int) int {

	freq := make(map[string]int)

	rev := func(s string) string {
		sb := []byte(s)
		slices.Reverse(sb)
		return strings.TrimLeft(string(sb), "0")
	}

	for _, num := range nums {
		s1 := strconv.Itoa(num)
		s2 := rev(s1)
		freq[s1]++
		freq[s2]++
	}

	return len(freq)
}
