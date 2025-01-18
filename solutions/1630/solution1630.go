package solution1630

import (
	"slices"
)

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1630
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_checkArithmeticSubarrays
	Benchmark_checkArithmeticSubarrays-24    	20167892	        64.02 ns/op	      51 B/op	       2 allocs/op
	PASS
*/

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {

	isValid := func(s []int) bool {
		if len(s) == 2 {
			return true
		}

		d := s[1] - s[0]
		for i := 1; i < len(s)-1; i++ {
			if s[i+1]-s[i] != d {
				return false
			}
		}

		return true
	}

	s := make([]int, 0, len(nums))
	output := make([]bool, len(l))
	for i := 0; i < len(l); i++ {
		a := l[i]
		b := r[i] + 1
		s = s[:b-a]
		copy(s, nums[a:b])
		slices.Sort(s)
		output[i] = isValid(s)
	}

	return output
}
