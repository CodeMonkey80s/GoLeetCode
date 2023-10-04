package solution401

import (
	"math"
	"math/bits"
	"strconv"
)

// ============================================================================
// 401. Binary Watch
// URL: https://leetcode.com/problems/binary-watch/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_readBinaryWatch-24    	   90870	     17138 ns/op	     608 B/op	      33 allocs/op
	PASS

*/

func readBinaryWatch(turnedOn int) []string {
	if turnedOn == 0 {
		return []string{"0:00"}
	}
	ans := []string{}
	for i := 1; i <= int(math.Pow(2, 10)); i++ {
		bc := bits.OnesCount(uint(i))
		if bc == turnedOn {
			t := ""
			hour := i >> 6
			if hour > 11 {
				continue
			}
			t += strconv.Itoa(hour) + ":"
			minute := 0b111111 & i
			if minute > 59 {
				continue
			}
			m := strconv.Itoa(minute)
			if len(m) < 2 {
				t += "0"
			}
			t += m
			ans = append(ans, t)
		}
	}
	return ans
}
