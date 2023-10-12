package solution1880

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
	val := func(s string) int {
		n := 0
		p := 1.0
		for i := len(s) - 1; i >= 0; i-- {
			p *= 10
			v := float64(byte(s[i]) - 'a')
			n += int(v * p)
		}
		return n
	}
	return val(firstWord)+val(secondWord) == val(targetWord)
}

func isSumEqual_array(firstWord string, secondWord string, targetWord string) bool {
	p := [8]int{
		1, 10, 100, 1_000, 10_000, 100_000, 1_000_000, 10_000_000,
	}
	n := 0
	i := 0
	v := 0
	val := func(s string) int {
		n = 0
		for i = len(s) - 1; i >= 0; i-- {
			v = int(byte(s[i]) - 'a')
			n += v * p[i]
		}
		return n
	}
	return val(firstWord)+val(secondWord) == val(targetWord)
}
