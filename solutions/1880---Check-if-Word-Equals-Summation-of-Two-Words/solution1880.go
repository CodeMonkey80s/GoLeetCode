package solution1880

import "math"

// ============================================================================
// 1880. Check if Word Equals Summation of Two Words
// URL: https://leetcode.com/problems/check-if-word-equals-summation-of-two-words/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1880---Check-if-Word-Equals-Summation-of-Two-Words
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_isSumEqual-24    	100000000	        10.55 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	getVal := func(s string) int {
		n := 0
		step := 0
		for i := len(s) - 1; i >= 0; i-- {
			p := math.Pow10(step)
			v := float64(byte(s[i]) - 'a')
			n += int(v * p)
			step++
		}
		return n
	}
	v1 := getVal(firstWord)
	v2 := getVal(secondWord)
	v3 := getVal(targetWord)
	return v1+v2 == v3
}
