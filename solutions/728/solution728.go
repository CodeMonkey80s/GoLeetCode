package solution728

import "strconv"

// ============================================================================
// 728. Self Dividing Numbers
// URL: https://leetcode.com/problems/self-dividing-numbers/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/728
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_selfDividingNumbers
	Benchmark_selfDividingNumbers-24    	 7992819	       149.7 ns/op	     248 B/op	       5 allocs/op
	PASS
*/

func selfDividingNumbers(left int, right int) []int {
	var numbers []int

outer:
	for i := left; i <= right; i++ {

		s := strconv.Itoa(i)
		for j := 0; j < len(s); j++ {
			switch {
			case s[j] == '0':
				continue outer
			case i%int(s[j]-'0') != 0:
				continue outer
			}
		}

		numbers = append(numbers, i)
	}

	return numbers
}
